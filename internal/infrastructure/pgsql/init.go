package pgsql

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func Init() (*sqlx.DB, error) {
	dbUser := os.Getenv("DATABASE_USER")
	dbName := os.Getenv("DATABASE_Name")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")

	if dbUser == "" || dbPass == "" || dbPort == "" || dbName == "" || dbHost == "" {
		return nil, fmt.Errorf("missing required database environment")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbPort, dbHost)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres %w", err)
	}

	log.Println("Postgresql connected suscesfully")
	return db, nil
}
