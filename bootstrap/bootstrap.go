package bootstrap

import (
	"beer-production-api/adapters/db/postgres"
	"database/sql"
)

type App struct {
	DB                *sql.DB
	BatchRepo   *postgres.BatchRepository
	BreweryRepo      *postgres.BreweryRepository
	EquipamentRepo *postgres.EquipamentRepository
	RecipeRepo          *postgres.RecipeRepository
	UserRepo     *postgres.UserRepository
}

func NewApp(db *sql.DB) *App {
	return &App{
		DB:                db,
		BatchRepo:   postgres.NewBatchRepository(db),
		BreweryRepo:      postgres.NewBreweryRepository(db),
		EquipamentRepo: postgres.NewEquipamentRepository(db),
		RecipeRepo:          postgres.NewRecipeRepository(db),
		UserRepo:     postgres.NewUserRepository(db),
	}
}