package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/brewery"
	breweryService "beer-production-api/service/brewery"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BreweryController struct {
	app *bootstrap.App
}

func NewBreweryController(app *bootstrap.App) *BreweryController {
	return &BreweryController{app: app}
}

func (bc *BreweryController) CreateBrewery(c echo.Context) (error) {
	fmt.Println("chegou no brewery controller")
	breweryDto := &brewery.CreateBreweryInputDto{}
	err := c.Bind(breweryDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	fmt.Printf("breweryDto: %+v", breweryDto)

	serviceInjection := breweryService.NewCreateBrewery(bc.app.BreweryRepo)
	output, err := serviceInjection.Execute(*breweryDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}