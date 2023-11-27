package RecipeService

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/recipe"
	"fmt"
)

type GetRecipeSteps struct {
	RecipeRepository recipe.RecipeRepository
	Breweryrepository brewery.BreweryRepository
}

func NewGetRecipeSteps(RecipeRepository recipe.RecipeRepository, BreweryRepository brewery.BreweryRepository) *GetRecipeSteps {
	return &GetRecipeSteps{
		RecipeRepository: RecipeRepository,
		Breweryrepository: BreweryRepository,
	}
}

func (gR *GetRecipeSteps) Execute(recipeId string, userId string, breweryId string) ([]*recipe.Steps, error) {
	role, err := gR.Breweryrepository.GetUserRole(userId, breweryId)
	if err != nil {
		return nil, err
	}

	if role == "" {
		return nil, fmt.Errorf("User doesn't have access to this brewery's recipes")
	}

	steps, err := gR.RecipeRepository.GetRecipeSteps(recipeId)
	if err != nil {
		return nil, err
	}

	return steps, nil
}