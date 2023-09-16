package brewery

import "fmt"

type Brewery struct {
	ID string
	Name string
	Cnpj string
	CreatorId string
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

	if t.CreatorId == "" {
		return fmt.Errorf("Creator id is required")
	}

	return nil
}