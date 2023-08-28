package batch

import (
	"fmt"
	"time"
)

type Batch struct {
	ID int
	Name string
	RecipeId int
	StartDate time.Time
	FinishDate time.Time
}

type BatchStep struct {
	ID int
	UserId int
	RecipeStepId int
	BatchId int
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

	if t.RecipeId <= 0 {
		return fmt.Errorf("Batch must have a recipe")
	}

	if t.StartDate.IsZero() {
		return fmt.Errorf("Batch must have a start date")
	}

	return nil
}