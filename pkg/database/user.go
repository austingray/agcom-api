package database

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Email          string `json:"email" validate:"required,email"`
	EmailValidated bool
	Handle         string
	Password       string `json:"password"`
	Hashed         string
}

// CreateUser handles creation of user in database
func (d *Database) CreateUser(email, hash string) bool {
	user := User{Email: email, Hashed: hash}

	// validation
	err := validate.Struct(user)
	if err != nil {
		// TODO: error logging/handling
		return false
	}

	// create the user
	if dbc := d.db.Create(&user); dbc.Error != nil {
		// TODO: error logging/handling
		return false
	}

	// success
	return true
}
