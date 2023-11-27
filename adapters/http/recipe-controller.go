package http

import (
	"beer-production-api/bootstrap"
	"beer-production-api/entities/recipe"
	recipeService "beer-production-api/service/recipe"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeController struct {
	app *bootstrap.App
}

func NewRecipeController(app *bootstrap.App) *RecipeController {
	return &RecipeController{app: app}
}

// @Summary Create recipe
// @Description Create recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param recipe body recipe.CreateRecipeInputDto true "recipe"
// @Success 200 {object} recipe.CreateRecipeInputDto
// @Failure 400 {object} common.HttpErrorResponse
// @Router /recipes [post]
func (rc *RecipeController) CreateRecipe(c echo.Context) (error) {
	userId := c.Get("userId")
	RecipeDto := &recipe.CreateRecipeInputDto{
		UserId: fmt.Sprintf("%v", userId),
	}
	err := c.Bind(RecipeDto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	serviceInjection := recipeService.NewCreateRecipe(rc.app.RecipeRepo, rc.app.BreweryRepo)
	output, err := serviceInjection.Execute(*RecipeDto)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}

// @Summary Get recipe by brewery id
// @Description Get a list of recipes from a brewery
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param brewery_id path string true "Brewery id"
// @Success 200 {object} recipe.GetRecipesByBreweryIdOutputDto
// @Failure 400 {object} common.HttpErrorResponse
// @Router /recipe/brewery/brewery-id [get]
func (rc *RecipeController) GetRecipesByBreweryId(c echo.Context) (error) {
	userId := c.Get("userId")
	var breweryId recipe.GetRecipesByBreweryIdInputDto
	err := c.Bind(&breweryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	serviceInjection := recipeService.NewGetRecipe(rc.app.RecipeRepo, rc.app.BreweryRepo)
	output, err := serviceInjection.Execute(breweryId.BreweryId, fmt.Sprintf("%v", userId))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, output)
}

// @Summary Get steps by recipe id
// @Description Get a list of steps from a recipe
// @Tags recipe
// @Accept  json
// @Produce  json
// @Param brewery_id path string true "Brewery id"
// @Param recipe_id path string true "Recipe id"
// @Success 200 {object} recipe.GetRecipesByBreweryIdOutputDto
// @Failure 400 {object} common.HttpErrorResponse
// @Router /recipe/brewery/brewery-id [get]
func (rc *RecipeController) GetRecipeSteps(c echo.Context) (error) {
	userId := c.Get("userId")

	var input recipe.GetRecipeStepsInputDto

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	serviceInjection := recipeService.NewGetRecipeSteps(rc.app.RecipeRepo, rc.app.BreweryRepo)

	output, err := serviceInjection.Execute(input.RecipeId, fmt.Sprintf("%v", userId), input.BreweryId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, output)
}