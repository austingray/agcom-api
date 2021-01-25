package server

import (
	"github.com/austingray/agcom-api/api"
	"github.com/austingray/agcom-api/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server engine and database
type Server struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

var server = Server{}

// Start starts the engine.
func Start() {
	server.Engine = gin.Default()
	server.DB = database.Default()

	api.RegisterRoutes(server.Engine)

	server.Engine.Run("0.0.0.0:8080")
}
