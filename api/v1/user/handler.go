package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type formData struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Register POST handler
func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// password to bytes
	b := []byte(password)

	// bcrypt
	p, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		// TODO: handle password error
		log.Println(err)
	}

	hashed := string(p)

	// TODO: create user in database

	log.Println(email + " : " + hashed)

	c.JSON(http.StatusOK, gin.H{"response": c.PostForm("stuff")})
}

// Login POST handler
func Login(c *gin.Context) {

}

// User GET handler
func User(c *gin.Context) {

}
