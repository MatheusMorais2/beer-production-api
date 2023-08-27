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

func NewBatchApplication() *Batch {
	return &Batch{}
}

func (t *Batch) isValid() error {
	if t.ID <= 0 {
		return fmt.Errorf("Wrong ID format")
	}

	if t.Name == "" {
		return fmt.Errorf("Batch name is required")
	}

	if t.RecipeId <= 0 {
		return fmt.Errorf("Batch must have a recipe")
	}

	if t.StartDate.IsZero() {
		return fmt.Errorf("Batch must have a start date")
	}

	if !t.FinishDate.IsZero() {
		return fmt.Errorf("Batch must not have a fisnish time when created")
	}

	return nil
}