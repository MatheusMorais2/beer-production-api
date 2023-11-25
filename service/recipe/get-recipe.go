package RecipeService

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/recipe"
	"fmt"
)

type GetRecipe struct {
	RecipeRepository recipe.RecipeRepository
	BreweryRepository brewery.BreweryRepository
}

func NewGetRecipe(RecipeRepository recipe.RecipeRepository, Breweryrepository brewery.BreweryRepository) *GetRecipe {
	return &GetRecipe{
		RecipeRepository: RecipeRepository,
		BreweryRepository: Breweryrepository,
	}
}

func (gR *GetRecipe) Execute(breweryId string, userId string) ([]*recipe.Recipe, error) {
	role, err := gR.BreweryRepository.GetUserRole(userId, breweryId)
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if role == "" {
		return nil, fmt.Errorf("User doesn't have access to this brewery's recipes")
	}

	recipes, err := gR.RecipeRepository.GetRecipesByBreweryId(breweryId)
	if err != nil {
		return nil, err
	}
	
	return recipes, nil
}