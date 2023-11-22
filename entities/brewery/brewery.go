package brewery

import "fmt"

type Brewery struct {
	ID string
	Name string
	Email string
}

type Invite struct {
	Id string
	InvitingUserId string
	InvitedUserId string
	BreweryId string
	Role string
}

func NewBreweryApplication() *Brewery {
	return &Brewery{}
}

func (t *Brewery) IsValid() error {
	if t.Name == "" {
		return fmt.Errorf("Brewery name is required")
	}

	if t.Email == "" {
		return fmt.Errorf("Brewery cnpj is required")
	}

	return nil
}