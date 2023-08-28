package postgres

import (
	"beer-production-api/entities/recipe"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func BuildInsertRecipeStepsQuery(recipeId int, steps []recipe.Steps) string {
	var sb strings.Builder
	sb.WriteString(`INSERT INTO "recipe_step" ("name", "recipe_id", "instruction")
	VALUES`)

	fmt.Println(strconv.Itoa(recipeId))
	for i := 0; i < len(steps); i++ {
		sb.WriteString(` ('` + steps[i].Name + `', ` + strconv.Itoa(recipeId) + `, '` + steps[i].Instruction + `')`)
		if i < (len(steps) - 1) {
			sb.WriteString(`,`)
		}
	}
	sb.WriteString(` RETURNING *;`)
	return sb.String()
}

func (t *RecipeRepository) Insert(recipeToInsert *recipe.Recipe) (*recipe.Recipe, error) {
	err := t.db.QueryRow(
		`INSERT INTO "recipe" (name, brewery_id)
		VALUES ($1, $2) RETURNING id`,
		recipeToInsert.Name, recipeToInsert.BreweryId,
	).Scan(&recipeToInsert.ID)
	if err != nil {
		fmt.Println("Error insert Recipe")
		return nil, err
	}

	recipeStepsQuery := BuildInsertRecipeStepsQuery(recipeToInsert.ID, recipeToInsert.Steps)
	fmt.Printf("recipe steps: %+v\n", recipeStepsQuery)
	rows, err := t.db.Query(recipeStepsQuery)
	if err != nil {
		fmt.Printf("error insert recipe steps: %+v", err.Error())
		fmt.Println("Error inserting recipe steps")
		return nil, err
	}
	fmt.Printf("rows: %+v", rows)

	return recipeToInsert, nil
}