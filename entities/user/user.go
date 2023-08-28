package user

import (
	"fmt"
)

type User struct {
	ID int
	Name string
	Role Role
	BreweryId int
}

type Role string

const (
	ReadOnly Role = "read-only"
	Technician Role = "technician"
	Admin Role = "admin"
	Unknown Role = "unknown"
)

func NewUserApplication() *User {
	return &User{}
}

func (t *User) IsValid() error {
	if t.Name == "" {
	fmt.Println("entrou no isvalid 2")

		return fmt.Errorf("User name is required")
	}

	if t.Role == Unknown {
	fmt.Println("entrou no isvalid 3")

		return fmt.Errorf("User must have a role")
	}

	if t.BreweryId <= 0 {
	fmt.Println("entrou no isvalid 4")

		return fmt.Errorf("User must be linked to a brewery")
	}
	
	return nil
}