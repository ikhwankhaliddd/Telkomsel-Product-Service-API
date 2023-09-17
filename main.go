package main

import (
	"fmt"
	"log"
	"os"

	categoryUsecase "github.com/ikhwankhaliddd/product-service/internal/components/category"
	categoryRepo "github.com/ikhwankhaliddd/product-service/internal/components/category/domain/repo"
	categoryPublicRepo "github.com/ikhwankhaliddd/product-service/internal/components/category/public_repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/products"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/domain/repo"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety"
	varietyRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/domain/repo"
	varietyPublicRepo "github.com/ikhwankhaliddd/product-service/internal/components/variety/public_repo"
	"github.com/ikhwankhaliddd/product-service/internal/helper/uploader"
	"github.com/joho/godotenv"

	categoryHandler "github.com/ikhwankhaliddd/product-service/internal/http/handler/category"
	productsHandler "github.com/ikhwankhaliddd/product-service/internal/http/handler/products"
	varietyHandler "github.com/ikhwankhaliddd/product-service/internal/http/handler/variety"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service Run on DEBUG mode")
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	createProductRepo := repo.NewInsertProduct(db)
	createVarietyPublicRepo := varietyPublicRepo.NewInsertVariety(db)
	getCategoryPublicRepo := categoryPublicRepo.NewCategoryGetter(db)
	createProductUsecase := products.NewProductCreator(createProductRepo, getCategoryPublicRepo, createVarietyPublicRepo)

	getProductRepo := repo.NewProductGetter(db)
	getVarietyPublicRepo := varietyPublicRepo.NewVarietyGetter(db)

	getProductUsecase := products.NewProductGetter(getProductRepo, getVarietyPublicRepo)
	getCategoryByIDRepo := categoryPublicRepo.NewCategoryByIDGetter(db)
	getCategoryUsecase := categoryUsecase.NewCategoryGetter(getCategoryByIDRepo)

	getProductListRepo := repo.NewProductListGetter(db)
	getProductListUsecase := products.NewProductListGetter(getProductListRepo, getVarietyPublicRepo, getCategoryByIDRepo)

	updateProductRepo := repo.NewProductUpdatter(db)
	updateProductUsecase := products.NewProductUpdatter(updateProductRepo)

	deleteProductRepo := repo.NewDeleteProduct(db)
	deleteProductUsecase := products.NewDeleteProduct(deleteProductRepo)

	productRatingRepo := repo.NewInsertRating(db)
	productRatingUsecase := products.NewPostRating(productRatingRepo)

	productsController := productsHandler.NewProductController(
		createProductUsecase,
		getProductUsecase,
		getProductListUsecase,
		updateProductUsecase,
		deleteProductUsecase,
		productRatingUsecase,
		getCategoryUsecase,
	)

	uploadImageRepo := varietyRepo.NewImageUploader(db)
	uploadHelper, _ := uploader.NewUploadHelper(viper.GetString("region"))
	uploadImageUsecase := variety.NewImageUploader(uploadImageRepo, uploadHelper)

	updateVarietyRepo := varietyRepo.NewVarietyUpdater(db)
	updateVarietyUsecase := variety.NewVarietyUpdatter(updateVarietyRepo)

	deleteVarietyRepo := varietyRepo.NewDeleteVariety(db)
	deleteVarietyUsecase := variety.NewDeleteVariety(deleteVarietyRepo)

	varietyController := varietyHandler.NewVarietyController(
		uploadImageUsecase,
		updateVarietyUsecase,
		deleteVarietyUsecase,
	)

	categoryCreatorRepo := categoryRepo.NewInsertCategory(db)
	categoryCreatorUsecase := categoryUsecase.NewCreateCategory(categoryCreatorRepo)

	categoryListGetterRepo := categoryRepo.NewCategoryListGetter(db)
	categoryListGetterUsecase := categoryUsecase.NewCategoryListGetter(categoryListGetterRepo)

	categoryController := categoryHandler.NewCategoryController(
		categoryCreatorUsecase,
		categoryListGetterUsecase,
	)
	v1 := e.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", productsController.HandleCreation)
			products.GET("/:id", productsController.HandleGetter)
			products.GET("", productsController.HandleListGetter)
			products.PATCH("/:id", productsController.HandleUpdatter)
			products.DELETE("/:id", productsController.HandleDelete)
			products.POST("/rating/:id", productsController.HandlePostRating)
		}
		variety := v1.Group("/variety")
		{
			variety.POST("/upload/:id", varietyController.HandleImageUpload)
			variety.PATCH("/:id", varietyController.HandleUpdatter)
			variety.DELETE("/:id", varietyController.HandleDelete)
		}
		category := v1.Group("/category")
		{
			category.POST("", categoryController.HandleCategoryCreator)
			category.GET("", categoryController.HandleCategoryListGetter)
		}
	}

	log.Fatal(e.Start(viper.GetString("server.address")))
}
