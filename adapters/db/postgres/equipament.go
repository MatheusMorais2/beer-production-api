package postgres

import (
	"beer-production-api/entities/equipament"
	"database/sql"
	"fmt"
)

type EquipamentRepository struct {
	db *sql.DB
}

func NewEquipamentRepository(db *sql.DB) *EquipamentRepository {
	return &EquipamentRepository{db: db}
}

func (t *EquipamentRepository) Insert(EquipamentToInsert *equipament.Equipament) (*equipament.Equipament, error) {
	err := t.db.QueryRow(
		`INSERT INTO "tblEquipament" (name)
		VALUES ($1) RETURNING id`,
		EquipamentToInsert.Name,
	).Scan(&EquipamentToInsert.ID)

	if err != nil {
		fmt.Println("Error on insert Equipament")
		return nil, err
	}

	return EquipamentToInsert, nil
}