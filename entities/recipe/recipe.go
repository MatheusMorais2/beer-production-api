package recipe

import (
	"fmt"
)

type Recipe struct {
	ID string
	Name string
	BreweryId string
	Steps []Steps
}

type Steps struct {
	Name string `json:"name"`
	Instruction string `json:"instruction"`
}

func NewRecipeApplication() *Recipe {
	return &Recipe{}
}

func (t *Recipe) IsValid() error {
	if t.Name == "" {
		return fmt.Errorf("Recipe name is required")
	}

	if t.BreweryId == "" {
		return fmt.Errorf("Recipe must be linked to a brewery")
	}

	if len(t.Steps) == 0 {
		return fmt.Errorf("Recipe must have steps")
	}

	return nil
}