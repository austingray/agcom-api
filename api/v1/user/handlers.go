package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// FormData for user POST requests
type FormData struct {
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
	log.Println(hashed)

	c.JSON(http.StatusOK, gin.H{"email": email})
}

// Login POST handler
func Login(c *gin.Context) {

}

// User GET handler
func User(c *gin.Context) {

}
