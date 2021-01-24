package user

import "github.com/gin-gonic/gin"

// RegisterRoutes registers api routes
func RegisterRoutes(e *gin.Engine) {
	group := e.Group("/api/v1/user")

	{
		group.POST("/register", Register)
		group.POST("/login", Login)
	}
}
