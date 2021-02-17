package user

import (
	"net/http"
	"os"
	"time"

	"github.com/austingray/agcom-api/pkg/database"
	"github.com/dgrijalva/jwt-go"
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

	// send registration email if client sets this value
	sendEmail := false
	if c.PostForm("sendEmail") != "" {
		sendEmail = true
	}

	d := c.MustGet("d").(*database.Database)
	user, err := d.CreateUser(email, password, sendEmail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Login POST handler
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	d := c.MustGet("d").(*database.Database)
	user, tx := d.GetUserByEmail(email)
	if tx.Error != nil {
		// if user not found
		// TODO: error logging and ambiguate error message
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// validate
	hp := []byte(user.Hashed)
	p := []byte(password)
	err := bcrypt.CompareHashAndPassword(hp, p)
	if err != nil {
		// incorrect password
		// TODO: error logging and ambiguate error message
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	// create token
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["id"] = user.ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		// jwt error
		// TODO: error logging and ambiguate error message
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// User GET handler
func User(c *gin.Context) {

}
