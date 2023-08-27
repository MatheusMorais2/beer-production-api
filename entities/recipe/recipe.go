package recipe

import (
	"fmt"
)

type Recipe struct {
	ID int
	Name string
	BreweryId int
}

func NewRecipeApplication() *Recipe {
	return &Recipe{}
}

func (t *Recipe) isValid() error {
	if t.ID <= 0 {
		return fmt.Errorf("Wrong ID format")
	}

	if t.Name == "" {
		return fmt.Errorf("Recipe name is required")
	}

	if t.BreweryId <= 0 {
		return fmt.Errorf("Recipe must be linked to a brewery")
	}

	return nil
}