package RecipeService

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/recipe"
	"fmt"
)

type CreateRecipe struct {
	RecipeRepository recipe.RecipeRepository
	BreweryRepository brewery.BreweryRepository
}

func NewCreateRecipe(RecipeRepository recipe.RecipeRepository, BreweryRepository brewery.BreweryRepository) *CreateRecipe {
	return &CreateRecipe{
		RecipeRepository: RecipeRepository,
		BreweryRepository: BreweryRepository,
	}
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

	role, err := cB.BreweryRepository.GetUserRole(input.UserId, input.BreweryId)
	if err != nil {
		return &recipe.CreateRecipeOutputDto{
			ErrorMessage: "Erro obtendo a função do usuário",
		}, err
	}

	if role != "admin" && role != "creator" && role != "brewer" {
		err = fmt.Errorf("Usuário sem permissão para criar receitas")
		return &recipe.CreateRecipeOutputDto{
			ErrorMessage: err.Error(),
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