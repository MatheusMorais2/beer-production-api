package postgres

import (
	"beer-production-api/entities/user"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (t *UserRepository) Insert(userToInsert *user.User) (*user.User, error) {
	fmt.Printf("userToInsert: %+v", userToInsert)
	err := t.db.QueryRow(
		`INSERT INTO "users" (name, password, email)
		VALUES ($1, $2, $3) RETURNING id`,
		userToInsert.Name,  userToInsert.Password, userToInsert.Email,
	).Scan(&userToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert User")
		return nil, err
	}

	return userToInsert, nil
}

func (t *UserRepository) GetByEmail(email string) (*user.User, error) {
	fmt.Printf("email no GetByEmail: %+v", email)
	var user *user.User
	err := t.db.QueryRow(
		`SELECT * FROM users WHERE email = $1`,
		email,
	).Scan(user)

	fmt.Printf("user no GetByEmail: %+v", user)
	if err != nil {
		fmt.Println("Error on insert User")
		return nil, err
	}

	return user, nil
}