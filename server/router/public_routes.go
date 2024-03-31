package router

import (
	"context"
	"github.com/gin-gonic/gin"
	userController "github.com/projects/watch-list/server/core/users"
	"github.com/projects/watch-list/server/database"
	"github.com/projects/watch-list/server/pkg/jwt/private"
)

func PublicRoutes(ctx context.Context, s *Server) (err error) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userCtrl, err := userController.Wire(database.GetCluster().Cluster)
	if err != nil {
		return
	}

	// User routes
	users := s.Engine.Group("api/v1/users")
	{
		users.GET("/me", private.AuthenticateJWT(), userCtrl.GetMyUser)
		users.GET("/:user_id", private.AuthenticateJWT(), userCtrl.GetUser)
		users.GET("", userCtrl.GetAllUsers)
	}
	return
}
