package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Default returns a default postgres connection using GORM.
func Default() *gorm.DB {
	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	database := os.Getenv("DATABASE")
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", hostname, port, username, database, password)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	} else {
		fmt.Println("Database connection successful.")
	}

	return db
}
