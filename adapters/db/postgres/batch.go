package postgres

import (
	"beer-production-api/entities/batch"
	"database/sql"
	"fmt"
)

type BatchRepository struct {
	db *sql.DB
}

func NewBatchRepository(db *sql.DB) *BatchRepository {
	return &BatchRepository{db: db}
}

func (t *BatchRepository) Insert(batchToInsert *batch.Batch) (*batch.Batch, error) {
	formatedStartDate := batchToInsert.StartDate.Format("2006-01-02 15:04:05")
	err := t.db.QueryRow(
		`INSERT INTO "batch" (name, recipe_id, start_date, finish_date)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		batchToInsert.Name, batchToInsert.RecipeId, formatedStartDate, nil,
	).Scan(&batchToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert Batch")
		return nil, err
	}

	return batchToInsert, nil
}

func (t *BatchRepository) InsertBatchStep(batchStepToInsert *batch.BatchStep) (*batch.BatchStep, error) {
	fmt.Printf("batchStepToInsert.BatchId: %+v\n", batchStepToInsert.BatchId)
	formatedStartDate := batchStepToInsert.StartedAt.Format("2006-01-02 15:04:05")
	err := t.db.QueryRow(
		`INSERT INTO "batch_recipe_step" (user_id, recipe_step_id, batch_id, started_at, finished_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, 
		batchStepToInsert.UserId, batchStepToInsert.RecipeStepId, batchStepToInsert.BatchId, formatedStartDate, nil,
	).Scan(&batchStepToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert batch step")
		return nil, err
	}

	return batchStepToInsert, nil
}