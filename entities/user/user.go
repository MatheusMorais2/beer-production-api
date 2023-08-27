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

func (t *User) isValid() error {
	if t.ID <= 0 {
		return fmt.Errorf("Wrong ID format")
	}

	if t.Name == "" {
		return fmt.Errorf("User name is required")
	}

	if t.Role == Unknown {
		return fmt.Errorf("User must have a role")
	}

	if t.BreweryId <= 0 {
		return fmt.Errorf("User must be linked to a brewery")
	}
	
	return nil
}