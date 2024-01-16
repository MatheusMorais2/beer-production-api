package main

import (
	"database/sql"
	"fmt"
	"os"

	"beer-production-api/adapters/db/postgres"
	server "beer-production-api/adapters/http"
	"beer-production-api/bootstrap"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectToDB() (*sql.DB, error) {
	godotenv.Load()
	dbConfig := postgres.Config{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
		Sslmode: os.Getenv("DB_SSL"),
	}

	return postgres.New(dbConfig)
}

func main() {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully connected to database")
	}
	defer db.Close()

	app := bootstrap.NewApp(db)

	httpServer := server.NewServer(app, db)
	httpServer.Start()
}