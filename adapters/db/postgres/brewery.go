package postgres

import (
	"beer-production-api/entities/brewery"
	"database/sql"
	"fmt"
)

type BreweryRepository struct {
	db *sql.DB
}

func NewBreweryRepository(db *sql.DB) *BreweryRepository {
	return &BreweryRepository{db: db}
}

func (t *BreweryRepository) Insert(breweryToInsert *brewery.Brewery) (*brewery.Brewery, error) {
	err := t.db.QueryRow(
		`INSERT INTO "tblBrewery" (name, cnpj)
		VALUES ($1, $2) RETURNING id`,
		breweryToInsert.Name, breweryToInsert.Cnpj,
	).Scan(&breweryToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert brewery")
		return nil, err
	}

	return breweryToInsert, nil
}