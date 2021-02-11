package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Email          string `json:"email" validate:"required,email"`
	EmailValidated bool
	Handle         string
	Hashed         string
}

// CreateUser handles creation of user in database
func (d *Database) CreateUser(email, password string) (User, error) {
	// password to bytes
	b := []byte(password)

	// bcrypt
	p, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		// TODO: handle password error
		log.Println(err)
	}

	// create user object
	user := User{Email: email, Hashed: string(p)}

	// validation
	err = validate.Struct(user)
	if err != nil {
		// TODO: error logging/handling
		log.Println("Error: User struct did not validate.")
		return user, err
	}

	// create the user
	if dbc := d.db.Create(&user); dbc.Error != nil {
		// TODO: error logging/handling
		log.Println("Error: Could not create user in db.")
		return user, err
	}

	// success
	return user, err
}

// GetUserByEmail d.db.Where("email = ?", email).First(&user)
func (d *Database) GetUserByEmail(email string) User {
	var user User
	d.db.Where("email = ?", email).First(&user)
	return user
}
