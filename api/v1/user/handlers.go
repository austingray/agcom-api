package user

import (
	"net/http"

	"github.com/austingray/agcom-api/pkg/database"
	"github.com/gin-gonic/gin"
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

	d := c.MustGet("d").(*database.Database)
	user, err := d.CreateUser(email, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Login POST handler
func Login(c *gin.Context) {

}

// User GET handler
func User(c *gin.Context) {

}
