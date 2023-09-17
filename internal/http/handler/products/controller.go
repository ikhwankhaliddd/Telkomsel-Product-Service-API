package products

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ikhwankhaliddd/product-service/internal/components/category"
	categoryValueType "github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/ikhwankhaliddd/product-service/internal/components/products/valuetype"
	varietyValueType "github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"

	"github.com/ikhwankhaliddd/product-service/internal/components/products"
	"github.com/ikhwankhaliddd/product-service/internal/helper/pagination"
	responder "github.com/ikhwankhaliddd/product-service/internal/helper/response"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productCreatorUsecase    products.ICreateProduct
	productGetterUsecase     products.IGetProduct
	productListGetterUsecase products.IGetListProduct
	productUpdatterUsecase   products.IUpdateProduct
	productDeleteUsecase     products.IDeleteProduct
	productRatingUsecase     products.IPostRating
	categoryUsecase          category.IGetCategoryByID
}

func NewProductController(
	productCreatorUsecase products.ICreateProduct,
	productGetterUsecase products.IGetProduct,
	productListGetterUsecase products.IGetListProduct,
	productUpdatterUsecase products.IUpdateProduct,
	productDeleteUsecase products.IDeleteProduct,
	productRatingUsecase products.IPostRating,
	categoryUsecase category.IGetCategoryByID,
) *productController {
	return &productController{
		productCreatorUsecase:    productCreatorUsecase,
		productGetterUsecase:     productGetterUsecase,
		productListGetterUsecase: productListGetterUsecase,
		productUpdatterUsecase:   productUpdatterUsecase,
		productDeleteUsecase:     productDeleteUsecase,
		productRatingUsecase:     productRatingUsecase,
		categoryUsecase:          categoryUsecase,
	}
}

func (controller *productController) HandleCreation(c echo.Context) error {
	ctx := c.Request().Context()
	type varietyRequest struct {
		VarietyName  string  `json:"variety_name" validate:"required"`
		VarietyPrice float64 `json:"variety_price" validate:"required"`
		VarietyStock int     `json:"variety_stock" validate:"required"`
	}

	type request struct {
		Name         string           `json:"name" validate:"required"`
		Description  string           `json:"description"`
		CategoryName string           `json:"category_name" validate:"required"`
		Variety      []varietyRequest `json:"variety" validate:"required"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	varietyData := []varietyValueType.CreateVarietyIn{}

	for _, val := range req.Variety {
		varietyData = append(varietyData, varietyValueType.CreateVarietyIn{
			Name:  val.VarietyName,
			Price: val.VarietyPrice,
			Stock: val.VarietyStock,
		})
	}

	input := valuetype.CreateProductIn{
		Name:        req.Name,
		Description: req.Description,
		Variety:     varietyData,
	}

	categoryInput := categoryValueType.GetCategoryIn{
		Name: req.CategoryName,
	}

	out, err := controller.productCreatorUsecase.CreateProduct(ctx, input, categoryInput)
	if err != nil {
		response := responder.CreateResponse("error create product", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responder.CreateResponse("Success Create Product", http.StatusOK, out.ID, nil)
	return c.JSON(http.StatusOK, response)
}

func (controller *productController) HandleGetter(c echo.Context) error {
	ctx := c.Request().Context()

	type varietyResponse struct {
		VarietyName  string  `json:"variety_name"`
		VarietyPrice float64 `json:"variety_price"`
		VarietyStock int     `json:"variety_stock"`
		VarietyImage string  `json:"variety_image"`
	}

	type response struct {
		ID           int               `json:"id"`
		Name         string            `json:"name"`
		Description  string            `json:"description"`
		Variety      []varietyResponse `json:"variety"`
		CategoryName string            `json:"category_name"`
		Rating       int               `json:"rating"`
		TotalStock   int               `json:"total_stock"`
		CreatedAt    time.Time         `json:"created_at"`
		UpdatedAt    time.Time         `json:"updated_at"`
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	product, err := controller.productGetterUsecase.GetByID(ctx, uint64(id))
	if err != nil {
		response := responder.CreateResponse("error get product", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	category, err := controller.categoryUsecase.GetByID(ctx, product.CategoryID)
	if err != nil {
		response := responder.CreateResponse("error get category", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	varietyResp := []varietyResponse{}
	totalStock := 0

	for _, val := range product.Varieties {
		totalStock += val.Stock
		varietyResp = append(varietyResp, varietyResponse{
			VarietyName:  val.Name,
			VarietyPrice: val.Price,
			VarietyStock: val.Stock,
			VarietyImage: val.Image,
		})
	}

	resp := response{
		ID:           int(product.ID),
		Name:         product.Name,
		Description:  product.Description,
		Variety:      varietyResp,
		CategoryName: category.Name,
		CreatedAt:    product.CreatedAt,
		Rating:       product.Rating,
		TotalStock:   totalStock,
		UpdatedAt:    product.UpdatedAt,
	}

	out := responder.CreateResponse("Success Get Product", http.StatusOK, resp, nil)
	return c.JSON(http.StatusOK, out)
}

func (controller *productController) HandleListGetter(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Page   int    `query:"page"`
		Limit  int    `query:"limit"`
		Search string `query:"search"`
	}

	type response struct {
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		Rating       int       `json:"rating"`
		CategoryName string    `json:"category_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	req := request{}
	resp := []response{}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	offset := pagination.CalculateOffset(req.Page, req.Limit)

	input := valuetype.GetProductListIn{
		Offset: offset,
		Limit:  req.Limit,
		Search: req.Search,
	}

	products, count, err := controller.productListGetterUsecase.GetListProduct(ctx, input)
	if err != nil {
		response := responder.CreateResponse("error get product list", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	for _, product := range products {
		category, err := controller.categoryUsecase.GetByID(ctx, product.CategoryID)
		if err != nil {
			response := responder.CreateResponse("error get category name", http.StatusInternalServerError, err.Error(), nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		resp = append(resp, response{
			ID:           int(product.ID),
			Name:         product.Name,
			Description:  product.Description,
			Rating:       product.Rating,
			CategoryName: category.Name,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		})
	}
	pagination := responder.PaginationObj{
		CurrentPage:   uint64(req.Page),
		RecordPerPage: uint64(input.Limit),
		Count:         uint64(count),
	}

	return c.JSON(http.StatusOK, responder.CreateResponse("success get product list", http.StatusOK, resp, &pagination))
}

func (controller *productController) HandleUpdatter(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	input := valuetype.UpdateProductIn{
		Name:        req.Name,
		Description: req.Description,
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = controller.productUpdatterUsecase.Update(ctx, uint64(reqID), input)
	if err != nil {
		response := responder.CreateResponse("error update product", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	out := responder.CreateResponse("Success Update Product", http.StatusOK, nil, nil)
	return c.JSON(http.StatusOK, out)
}

func (controller *productController) HandleDelete(c echo.Context) error {
	ctx := c.Request().Context()

	req := c.Param("id")

	if req == "" {
		return c.JSON(http.StatusBadRequest, responder.CreateResponse("invalid product id", http.StatusBadRequest, nil, nil))
	}

	reqInt, err := strconv.Atoi(req)
	if err != nil {
		return err
	}

	err = controller.productDeleteUsecase.DeleteProduct(ctx, uint64(reqInt))
	if err != nil {
		response := responder.CreateResponse("error delete product", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, responder.CreateResponse("success delete a product", http.StatusOK, nil, nil))
}

func (controller *productController) HandlePostRating(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Rating int `json:"rating"`
	}
	req := request{}

	reqID := c.Param("id")

	if reqID == "" {
		return c.JSON(http.StatusBadRequest, responder.CreateResponse("invalid product id", http.StatusBadRequest, nil, nil))
	}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	reqIDInt, err := strconv.Atoi(reqID)
	if err != nil {
		return err
	}

	input := valuetype.PostRatingIn{
		ID:     uint64(reqIDInt),
		Rating: req.Rating,
	}

	err = controller.productRatingUsecase.PostRating(ctx, input)
	if err != nil {
		response := responder.CreateResponse("error post product rating", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	out := responder.CreateResponse("Success Post Product Rating", http.StatusOK, nil, nil)
	return c.JSON(http.StatusOK, out)
}
