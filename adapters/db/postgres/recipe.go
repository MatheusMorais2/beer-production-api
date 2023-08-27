package postgres

import (
	"beer-production-api/entities/recipe"
	"database/sql"
	"fmt"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (t *RecipeRepository) Insert(RecipeToInsert *recipe.Recipe) (*recipe.Recipe, error) {
	err := t.db.QueryRow(
		`INSERT INTO "tblRecipe" (name, brewery_id)
		VALUES ($1, $2) RETURNING id`,
		RecipeToInsert.Name, RecipeToInsert.BreweryId,
	).Scan(&RecipeToInsert.ID)

	if err != nil {
		fmt.Println("Error insert Recipe")
		return nil, err
	}

	return RecipeToInsert, nil
}