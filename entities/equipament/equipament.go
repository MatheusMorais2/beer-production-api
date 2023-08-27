package equipament

import (
	"fmt"
)

type Equipament struct {
	ID int
	Name string
}

func NewEquipamentApplication() *Equipament {
	return &Equipament{}
}

func (t *Equipament) isValid() error {
	if t.ID <= 0 {
		return fmt.Errorf("Wrong ID format")
	}

	if t.Name == "" {
		return fmt.Errorf("Equipament name is required")
	}

	return nil
}