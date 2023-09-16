package RecipeService

import (
	"beer-production-api/entities/recipe"
)

type CreateRecipe struct {
	RecipeRepository recipe.RecipeRepository
}

func NewCreateRecipe(RecipeRepository recipe.RecipeRepository) *CreateRecipe {
	return &CreateRecipe{RecipeRepository: RecipeRepository}
}

func (cB *CreateRecipe) Execute(input recipe.CreateRecipeInputDto) (*recipe.CreateRecipeOutputDto, error) {
	RecipeToCreate := recipe.NewRecipeApplication()
	RecipeToCreate.Name = input.Name
	RecipeToCreate.BreweryId = input.BreweryId
	RecipeToCreate.Steps = input.Steps
	err := RecipeToCreate.IsValid()
	if err != nil {
		return &recipe.CreateRecipeOutputDto{
			ErrorMessage: "Receita inválida ou com informações faltando",
		}, err
	}

	createdRecipe, err := cB.RecipeRepository.Insert(RecipeToCreate)
	if err != nil {
		return &recipe.CreateRecipeOutputDto{
			ErrorMessage: "Erro na criação da receita no banco de dados",
		}, err
	}

	output := &recipe.CreateRecipeOutputDto{
		Id: createdRecipe.ID,
		Name: createdRecipe.Name,
		BreweryId: createdRecipe.BreweryId,
		Steps: createdRecipe.Steps,
	}

	return output, nil
}