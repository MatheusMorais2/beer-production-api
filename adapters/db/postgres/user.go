package postgres

import (
	"beer-production-api/entities/user"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (t *UserRepository) Insert(userToInsert *user.User) (*user.User, error) {
	err := t.db.QueryRow(
		`INSERT INTO "users" (name, password, email)
		VALUES ($1, $2, $3) RETURNING id`,
		userToInsert.Name,  userToInsert.Password, userToInsert.Email,
	).Scan(&userToInsert.ID)

	if err != nil {
		return nil, err
	}

	return userToInsert, nil
}

func (t *UserRepository) GetByEmail(email string) (*user.User, error) {
	user := &user.User{}
	err := t.db.QueryRow(
		`SELECT id, name, email, password FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}