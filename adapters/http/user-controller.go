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

// @Summary Create user
// @Description Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateUserInputDto true "user"
// @Success 200 {object} user.CreateUserInputDto
// @Failure 400 {object} common.HttpErrorResponse
// @Router /user [post]
func (uc *UserController) CreateUser(c echo.Context) (error) {
	userDto := &user.CreateUserInputDto{}
	err := c.Bind(userDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	serviceInjection := userService.NewCreateUser(uc.app.UserRepo)
	output, err := serviceInjection.Execute(*userDto)
	if err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, output)
}

func (uc *UserController) Login(c echo.Context) (error) {
	fmt.Println("alou")
	loginInputDto := &user.LoginUserInputDto{}
	err := c.Bind(loginInputDto)
	if err != nil {
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
