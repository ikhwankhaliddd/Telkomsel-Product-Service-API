package variety

import (
	"net/http"
	"strconv"

	"github.com/ikhwankhaliddd/product-service/internal/components/variety"
	"github.com/ikhwankhaliddd/product-service/internal/components/variety/valuetype"
	"github.com/labstack/echo/v4"

	responder "github.com/ikhwankhaliddd/product-service/internal/helper/response"
)

type varietyController struct {
	imageUploaderUsecase   variety.IUploadImage
	varietyUpdatterUsecase variety.IUpdateVariety
	deleteVarietyUsecase   variety.IDeleteVariety
}

func NewVarietyController(
	imageUploaderUsecase variety.IUploadImage,
	varietyUpdatterUsecase variety.IUpdateVariety,
	deleteVarietyUsecase variety.IDeleteVariety,
) *varietyController {
	return &varietyController{
		imageUploaderUsecase:   imageUploaderUsecase,
		varietyUpdatterUsecase: varietyUpdatterUsecase,
		deleteVarietyUsecase:   deleteVarietyUsecase,
	}
}

func (controller *varietyController) HandleImageUpload(c echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("file")
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responder.CreateResponse("invalid file upload", http.StatusBadRequest, nil, nil))
	}

	input := valuetype.UploadImageIn{
		ID:   idInt,
		File: file,
	}

	err = controller.imageUploaderUsecase.UploadImage(ctx, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responder.CreateResponse("failed upload image", http.StatusInternalServerError, err, nil))
	}

	return c.JSON(http.StatusOK, responder.CreateResponse("success upload image", http.StatusOK, nil, nil))
}

func (controller *varietyController) HandleUpdatter(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		response := responder.CreateResponse("error binding request", http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	input := valuetype.UpdateVarietyIn{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = controller.varietyUpdatterUsecase.Update(ctx, uint64(reqID), input)
	if err != nil {
		response := responder.CreateResponse("error update variety", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	out := responder.CreateResponse("Success Update variety", http.StatusOK, nil, nil)
	return c.JSON(http.StatusOK, out)
}

func (controller *varietyController) HandleDelete(c echo.Context) error {
	ctx := c.Request().Context()

	req := c.Param("id")

	if req == "" {
		return c.JSON(http.StatusBadRequest, responder.CreateResponse("invalid variety id", http.StatusBadRequest, nil, nil))
	}

	reqInt, err := strconv.Atoi(req)
	if err != nil {
		return err
	}

	err = controller.deleteVarietyUsecase.DeleteVariety(ctx, uint64(reqInt))
	if err != nil {
		response := responder.CreateResponse("error delete variety", http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, responder.CreateResponse("success delete a variety", http.StatusOK, nil, nil))
}
