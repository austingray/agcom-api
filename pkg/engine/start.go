package engine

import (
	"github.com/austingray/agcom-api/api/v1/user"
	"github.com/gin-gonic/gin"
)

var e *gin.Engine

// Start starts the engine.
func Start() {
	e = gin.Default()

	// register routes
	user.RegisterRoutes(e)

	e.Run("0.0.0.0:8080")
}
