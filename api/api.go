package api

import (
	"github.com/austingray/agcom-api/api/v1/user"
	"github.com/austingray/agcom-api/pkg/server"
)

// RegisterRoutes registers all API routes to the server in a single spot
func RegisterRoutes(s server.Server) {
	group := s.Engine.Group("/api/v1")

	// user
	{
		group.POST("/user/register", user.Register)
		group.POST("/user/login", user.Login)
	}
}
