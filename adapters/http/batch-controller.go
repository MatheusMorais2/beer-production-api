package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/batch"
	batchService "beer-production-api/service/batch"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BatchController struct {
	app *bootstrap.App
}

func NewBatchController(app *bootstrap.App) *BatchController {
	return &BatchController{app: app}
}

func (bc *BatchController) CreateBatch(c echo.Context) (error) {
	fmt.Println("chegou no Batch controller")
	batchDto := &batch.CreateBatchInputDto{}
	err := c.Bind(batchDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	fmt.Printf("BatchDto: %+v\n", batchDto)

	serviceInjection := batchService.NewCreateBatch(bc.app.BatchRepo)
	output, err := serviceInjection.Execute(*batchDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}

func (bc *BatchController) CreateBatchStep(c echo.Context) (error) {
	fmt.Println("chegou no Batch controller")
	batchStepDto := &batch.CreateBatchStepInputDto{}
	err := c.Bind(batchStepDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	fmt.Printf("BatchStepDto: %+v\n", batchStepDto)

	serviceInjection := batchService.NewCreateBatchStep(bc.app.BatchRepo)
	output, err := serviceInjection.Execute(*batchStepDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}