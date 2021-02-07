package server

import (
	"github.com/austingray/agcom-api/api"
	"github.com/austingray/agcom-api/pkg/database"
	"github.com/gin-gonic/gin"
)

// Server engine and database
type Server struct {
	Database *database.Database
	Engine   *gin.Engine
}

var server = Server{}

// Start starts the server.
func Start() {
	server.Engine = gin.Default()
	server.Database = database.Default()

	server.Engine.Use(func(c *gin.Context) {
		c.Set("d", server.Database)
		c.Next()
	})

	api.RegisterRoutes(server.Engine)

	server.Engine.Run("0.0.0.0:8080")
}
