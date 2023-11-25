package postgres

import (
	"beer-production-api/entities/recipe"
	"database/sql"
	"fmt"
	"strings"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func BuildInsertRecipeStepsQuery(recipeId string, steps []recipe.Steps) string {
	var sb strings.Builder
	sb.WriteString(`INSERT INTO "recipe_step" ("name", "recipe_id", "instruction")
	VALUES`)

	for i := 0; i < len(steps); i++ {
		sb.WriteString(` ('` + steps[i].Name + `', '` + recipeId + `', '` + steps[i].Instruction + `')`)
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
		VALUES ($1, $2) RETURNING id;`,
		recipeToInsert.Name, recipeToInsert.BreweryId,
	).Scan(&recipeToInsert.ID)
	if err != nil {
		return nil, err
	}

	recipeStepsQuery := BuildInsertRecipeStepsQuery(recipeToInsert.ID, recipeToInsert.Steps)
	_, err = t.db.Query(recipeStepsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return recipeToInsert, nil
}

// GET RECIPES BY BREWERY ID
func (t *RecipeRepository) GetRecipesByBreweryId(breweryId string) ([]*recipe.Recipe, error) {
    recipes := []*recipe.Recipe{}
    rows, err := t.db.Query(
        `SELECT r.id , r.name , r.brewery_id
		FROM recipe r
		WHERE r.brewery_id = $1`,
        breweryId,
    )

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        recipeForBrewery := &recipe.Recipe{}
        err := rows.Scan(&recipeForBrewery.ID, &recipeForBrewery.Name, &recipeForBrewery.BreweryId)
        if err != nil {
            return nil, err
        }
        recipes = append(recipes, recipeForBrewery)
    }

    return recipes, nil
}

func (t *RecipeRepository) GetRecipeSteps(recipeId string) ([]*recipe.Steps, error) {
	steps := []*recipe.Steps{}

	rows, err := t.db.Query(
		`
			SELECT name, instruction 
			FROM recipe_step 
			WHERE recipe_id = $1
		`, recipeId)

	if err != nil {
        return nil, err
    }

	defer rows.Close()

	for rows.Next() {
		step := &recipe.Steps{}
		err = rows.Scan(&step.Name, &step.Instruction)
		if err != nil {
            return nil, err
        }

		steps = append(steps, step)
	}

	return steps, nil
}