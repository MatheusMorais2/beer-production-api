package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/recipe"
	recipeService "beer-production-api/service/recipe"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeController struct {
	app *bootstrap.App
}

func NewRecipeController(app *bootstrap.App) *RecipeController {
	return &RecipeController{app: app}
}

func (rc *RecipeController) CreateRecipe(c echo.Context) (error) {
	RecipeDto := &recipe.CreateRecipeInputDto{}
	err := c.Bind(RecipeDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	serviceInjection := recipeService.NewCreateRecipe(rc.app.RecipeRepo)
	output, err := serviceInjection.Execute(*RecipeDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}