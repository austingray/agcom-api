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

// Default returns a default Database object with a connection to the dev database
func Default() *Database {
	database := os.Getenv("POSTGRES_DB")
	db := connect(database)
	d := postConnect(db)
	return d
}

// Test returns a test Database object with a connection to the test database
func Test() *Database {
	database := "testdb"
	db := connect(database)
	d := postConnect(db)
	return d
}

func connect(database string) *gorm.DB {
	hostname := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USER")
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

func postConnect(db *gorm.DB) *Database {
	var d = &Database{}
	d.db = db

	// migrations
	d.db.AutoMigrate(&User{})

	// init validator
	validate = validator.New()

	return d
}
