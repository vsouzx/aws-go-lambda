package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb() (*sqlx.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbEndpoint := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUsername, dbPassword, dbEndpoint, dbName)
	fmt.Println("Connecting to database with DSN:", dsn)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}else {
		fmt.Println("Database connection established successfully.")
	}

	return db, err
}
