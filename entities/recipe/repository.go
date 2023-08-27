package recipe

type RecipeRepository interface {
	Insert(recipe *Recipe) (*Recipe, error)
}