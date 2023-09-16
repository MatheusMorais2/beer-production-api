package equipament

import (
	"fmt"
)

type Equipament struct {
	ID string
	Name string
}

func NewEquipamentApplication() *Equipament {
	return &Equipament{}
}

func (t *Equipament) isValid() error {
	if t.ID == "" {
		return fmt.Errorf("Wrong ID format")
	}

	if t.Name == "" {
		return fmt.Errorf("Equipament name is required")
	}

	return nil
}