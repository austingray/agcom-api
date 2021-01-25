package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Default returns a default postgres connection using GORM.
func Default() *gorm.DB {
	hostname := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USER")
	database := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASS")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", hostname, port, username, database, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	} else {
		fmt.Println("Database connection successful.")
	}

	return db
}
