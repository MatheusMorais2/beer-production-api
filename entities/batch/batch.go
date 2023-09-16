package batch

import (
	"fmt"
	"time"
)

type Batch struct {
	ID string
	Name string
	RecipeId string
	StartDate time.Time
	FinishDate time.Time
}

type BatchStep struct {
	ID string
	UserId string
	RecipeStepId string
	BatchId string
	StartedAt time.Time
	FinishedAt time.Time
}

func NewBatchApplication() *Batch {
	return &Batch{}
}

func NewBatchStepApplication() *BatchStep {
	return &BatchStep{}
}

func (t *Batch) IsValid() error {

	if t.Name == "" {
		return fmt.Errorf("Batch name is required")
	}

	if t.RecipeId == "" {
		return fmt.Errorf("Batch must have a recipe")
	}

	if t.StartDate.IsZero() {
		return fmt.Errorf("Batch must have a start date")
	}

	return nil
}