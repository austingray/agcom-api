package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database struct for associating methods with gorm.DB
type Database struct {
	db *gorm.DB
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Default returns a default Database object with a gorm db instance
func Default() *Database {
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

	var d = &Database{}
	d.db = db

	// migrations
	d.db.AutoMigrate(&User{})

	// init validator
	validate = validator.New()

	return d
}
