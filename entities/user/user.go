package user

import (
	"fmt"
)

type User struct {
	ID int
	Name string
	Password string
	Email string
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

	if t.Password == "" {
		fmt.Println("entrou no isvalid user password")

		return fmt.Errorf("User must have a password") 
	}

	if t.Email == "" {
		fmt.Println("entrou no isvalid user email")

		return fmt.Errorf("User must have a email") 
	}
 	
	return nil
}