package brewery

import "fmt"

type Brewery struct {
	ID int
	Name string
	Cnpj string
}

func NewBreweryApplication() *Brewery {
	return &Brewery{}
}

func (t *Brewery) IsValid() error {

	if t.Name == "" {
		return fmt.Errorf("Brewery name is required")
	}

	if t.Cnpj == "" {
		return fmt.Errorf("Brewery cnpj is required")
	}

	return nil
}