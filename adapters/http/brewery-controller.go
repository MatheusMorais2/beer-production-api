package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/brewery"
	breweryService "beer-production-api/service/brewery"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BreweryController struct {
	app *bootstrap.App
}

func NewBreweryController(app *bootstrap.App) *BreweryController {
	return &BreweryController{app: app}
}

// @Summary Create brewery
// @Description Create brewery
// @Tags brewery
// @Accept  json
// @Produce  json
// @Param brewery body brewery.CreateBreweryInputDto true "brewery"
// @Success 200 {object} brewery.CreateBreweryInputDto
// @Failure 400 {object} common.HttpErrorResponse
// @Router /brewery [post]
func (bc *BreweryController) CreateBrewery(c echo.Context) (error) {
	breweryDto := &brewery.CreateBreweryInputDto{}
	err := c.Bind(breweryDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	serviceInjection := breweryService.NewCreateBrewery(bc.app.BreweryRepo)
	output, err := serviceInjection.Execute(*breweryDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}