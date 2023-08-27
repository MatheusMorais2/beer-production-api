package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	app *bootstrap.App
}

func NewUserController(app *bootstrap.App) *UserController {
	return &UserController{app: app}
}

func (uc *UserController) CreateUser(c echo.Context) (error) {
	userDto := user.CreateUserInputDto{}
	err := c.Bind(userDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	//serviceInjection := service.New
	return nil
}
