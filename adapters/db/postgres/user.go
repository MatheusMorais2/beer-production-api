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
	fmt.Printf("userToInsert: %+v\n", userToInsert)
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
	fmt.Printf("email no GetByEmail: %+v\n", email)
	user := &user.User{}
	err := t.db.QueryRow(
		`SELECT id, name, email, password FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	fmt.Printf("user no GetByEmail: %+v\n", user)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error on get by email User -- no rows: %+v", err)
			return nil, err
		}
		fmt.Printf("Error on get by email User: %+v", err)
		return nil, err
	}

	return user, nil
}