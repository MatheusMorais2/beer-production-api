package recipe

type RecipeRepository interface {
	Insert(recipe *Recipe) (*Recipe, error)
	GetRecipesByBreweryId(breweryId string) ([]*Recipe, error)
	GetRecipeSteps(recipeId string) ([]*Steps, error)
}