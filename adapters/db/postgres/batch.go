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
	err := t.db.QueryRow(
		`INSERT INTO "tblBatch" (name, recipe_id, start_date, finish_date)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		batchToInsert.Name, batchToInsert.RecipeId, batchToInsert.StartDate, batchToInsert.FinishDate,
	).Scan(&batchToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert Batch")
		return nil, err
	}

	return batchToInsert, nil
}