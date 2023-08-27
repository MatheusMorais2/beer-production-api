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

func (t *UserRepository) Insert(UserToInsert *user.User) (*user.User, error) {
	err := t.db.QueryRow(
		`INSERT INTO "tblUser" (name, role, brewery_id)
		VALUES ($1, $2, $3) RETURNING id`,
		UserToInsert.Name, UserToInsert.Role, UserToInsert.BreweryId,
	).Scan(&UserToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert User")
		return nil, err
	}

	return UserToInsert, nil
}