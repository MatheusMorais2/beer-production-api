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

	userId := c.Get("userId")

	breweryDto.UserId = fmt.Sprintf("%v", userId)

	serviceInjection := breweryService.NewCreateBrewery(bc.app.BreweryRepo, bc.app.UserRepo)
	output, err := serviceInjection.Execute(*breweryDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}

// @Summary Get brewery by user id
// @Description Get brewery by user id
// @Tags brewery
// @Accept  json
// @Produce  json
// @Param brewery path string true "User id"
// @Success 200 {object} brewery.GetUserBreweriesOutputDTO
// @Failure 400 {object} common.HttpErrorResponse
// @Router /brewery [get]
func (bc *BreweryController) GetBreweriesByUserId(c echo.Context) (error) {
	userId := c.Get("userId")

	serviceInjection := breweryService.NewGetBrewery(bc.app.BreweryRepo)
	output, err := serviceInjection.Execute(fmt.Sprintf("%v", userId))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}


// @Summary Invite user
// @Description Invite user to a brewery
// @Tags brewery
// @Accept  json
// @Produce  json
// @Param invite body brewery.InviteUserInputDTO true "invite"
// @Success 200 {object} brewery.InviteUserOutputDTO
// @Failure 400 {object} common.HttpErrorResponse
// @Router /brewery/invite [post]
func (bc *BreweryController) InviteUser(c echo.Context) (error) {
	invintingUserId := c.Get("userId")

	invite := brewery.InviteUserInputDTO{
		InvitingUserId: fmt.Sprintf("%v", invintingUserId),
	}

	err := c.Bind(&invite)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	serviceInjection := breweryService.NewInviteUserToBrewery(bc.app.BreweryRepo, bc.app.UserRepo)
	createdInvite, err := serviceInjection.Execute(invite)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, createdInvite)
}