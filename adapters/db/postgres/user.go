package postgres

import (
	"beer-production-api/common"
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
		return nil, common.NewErrorResponse(common.InternalError, err.Error())
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
			return nil, common.NewErrorResponse(common.NotFound, err.Error())
		}
		return nil, common.NewErrorResponse(common.InternalError, err.Error())
	}

	return user, nil
}

func (t *UserRepository) GetById(id string) (*user.User, error) {
	user := &user.User{}
	
	err := t.db.QueryRow(
		`SELECT id, name, email, password FROM users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.NewErrorResponse(common.NotFound, err.Error())
		}
		return nil, common.NewErrorResponse(common.InternalError, err.Error())
	}

	return user, nil
}