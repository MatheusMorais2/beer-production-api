package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/user"
	userService "beer-production-api/service/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	app *bootstrap.App
}

func NewUserController(app *bootstrap.App) *UserController {
	return &UserController{app: app}
}

func (uc *UserController) CreateUser(c echo.Context) (error) {
	fmt.Println("chegou no user controller")
	userDto := &user.CreateUserInputDto{}
	err := c.Bind(userDto)
	if err != nil {
		fmt.Println("entrou no erro do controller")
		return c.JSON(400, err.Error())
	}

	serviceInjection := userService.NewCreateUser(uc.app.UserRepo)
	output, err := serviceInjection.Execute(*userDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, output)
}

func (uc *UserController) Login(c echo.Context) (error) {
	fmt.Println("chegou no login controller")
	loginInputDto := &user.LoginUserInputDto{}
	err := c.Bind(loginInputDto)
	if err != nil {
		fmt.Println("entrou no erro do login controller")
		return c.JSON(400, err.Error())
	}

	serviceInjection := userService.NewLoginUser(uc.app.UserRepo)
	output, err := serviceInjection.Execute(*loginInputDto)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized, err)
	}
	return c.JSON(http.StatusAccepted, output)
}
