package postgres

import (
	"beer-production-api/entities/brewery"
	"database/sql"
)

type BreweryRepository struct {
	db *sql.DB
}

func NewBreweryRepository(db *sql.DB) *BreweryRepository {
	return &BreweryRepository{db: db}
}

func (t *BreweryRepository) Insert(breweryToInsert *brewery.Brewery) (*brewery.Brewery, error) {
	err := t.db.QueryRow(
		`INSERT INTO "brewery" (name, cnpj)
		VALUES ($1, $2) RETURNING id`,
		breweryToInsert.Name, breweryToInsert.Cnpj,
	).Scan(&breweryToInsert.ID)

	if err != nil {
		return nil, err
	}

	return breweryToInsert, nil
}