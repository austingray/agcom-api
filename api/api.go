package api

import (
	"github.com/austingray/agcom-api/api/v1/user"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes to the server in a single spot
func RegisterRoutes(e *gin.Engine) {
	group := e.Group("/api/v1")

	// user
	{
		group.POST("/user/register", user.Register)
		group.POST("/user/login", user.Login)
	}
}
