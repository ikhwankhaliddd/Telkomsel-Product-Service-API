package category

import (
	"net/http"

	"github.com/ikhwankhaliddd/product-service/internal/components/category"
	"github.com/ikhwankhaliddd/product-service/internal/components/category/valuetype"
	"github.com/labstack/echo/v4"

	responder "github.com/ikhwankhaliddd/product-service/internal/helper/response"
)

type categoryController struct {
	categoryCreatorUsecase    category.ICreateCategory
	categoryListGetterUsecase category.IGetCategoryList
}

func NewCategoryController(
	categoryCreatorUsecase category.ICreateCategory,
	categoryListGetterUsecase category.IGetCategoryList,
) *categoryController {
	return &categoryController{
		categoryCreatorUsecase:    categoryCreatorUsecase,
		categoryListGetterUsecase: categoryListGetterUsecase,
	}
}

func (controller *categoryController) HandleCategoryCreator(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Name string `json:"name"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	input := valuetype.InsertCategoryIn{
		Name: req.Name,
	}
	err := controller.categoryCreatorUsecase.CreateCategory(ctx, input)
	if err != nil {
		response := responder.CreateResponse("error create category", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responder.CreateResponse("Success Create Category", http.StatusOK, nil, nil)
	return c.JSON(http.StatusOK, response)
}

func (controller *categoryController) HandleCategoryListGetter(c echo.Context) error {
	ctx := c.Request().Context()

	type response struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	resp := []response{}
	categoryList, err := controller.categoryListGetterUsecase.GetCategoryList(ctx)
	if err != nil {
		response := responder.CreateResponse("error get category list", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	for _, val := range categoryList {
		resp = append(resp, response{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	out := responder.CreateResponse("Success Get Category List", http.StatusOK, resp, nil)
	return c.JSON(http.StatusOK, out)
}
