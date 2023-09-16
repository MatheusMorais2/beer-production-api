package user

import (
	"fmt"
)

type User struct {
	ID string
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
		return fmt.Errorf("User name is required")
	}

	if t.Password == "" {
		return fmt.Errorf("User must have a password") 
	}

	if t.Email == "" {
		return fmt.Errorf("User must have a email") 
	}
 	
	return nil
}