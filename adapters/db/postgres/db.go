package postgres

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	Sslmode string
}

func New(config Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.Sslmode))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}