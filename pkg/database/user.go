package database

import (
	"errors"
	"log"
	"unicode"

	"github.com/austingray/agcom-api/pkg/smtp"
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
func (d *Database) CreateUser(email, password string, sendEmail bool) (User, error) {
	// create the user and validate the email
	user := User{Email: email}
	err := validate.Struct(user)
	if err != nil {
		// TODO: error logging/handling
		log.Println("Error: User struct did not validate.")
		return user, err
	}

	// bail if existing user
	_, tx := d.GetUserByEmail(email)
	if tx.Error == nil {
		// TODO: error logging/handling
		log.Println("Error: User " + email + " already exists.")
		return user, errors.New("user " + email + " already exists")
	}

	// validate password
	if validatePassword(password) == false {
		// TODO: error logging/handling
		log.Println("Error: Password not complex enough.")
		return user, errors.New("password not complex enough")
	}

	// hash the password
	b := []byte(password)
	p, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		// TODO: handle bcrypt error
		log.Println(err)
	}
	user.Hashed = string(p)

	// create the user
	if dbc := d.db.Create(&user); dbc.Error != nil {
		// TODO: error logging/handling
		log.Println("Error: Could not create user in db.")
		return user, dbc.Error
	}

	if sendEmail == true {
		smtp.SendEmail(email, "registration")
	}

	// success
	return user, err
}

// GetUserByEmail d.db.Where("email = ?", email).First(&user)
func (d *Database) GetUserByEmail(email string) (User, *gorm.DB) {
	var user User
	tx := d.db.Where("email = ?", email).First(&user)
	return user, tx
}

// passwords must have a number, uppercase letter, special character, min length of 8
func validatePassword(s string) bool {
	var number, upper, special bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		default:
		}
	}

	return number && upper && special && len(s) > 7
}
