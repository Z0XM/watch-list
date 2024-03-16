package router

import (
	"context"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(ctx context.Context, s *Server) (err error) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return
}
