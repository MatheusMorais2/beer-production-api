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
		`INSERT INTO "users" (name, role, brewery_id)
		VALUES ($1, $2, $3) RETURNING id`,
		userToInsert.Name, userToInsert.Role, userToInsert.BreweryId,
	).Scan(&userToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert User")
		return nil, err
	}

	return userToInsert, nil
}