package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/projects/watch-list/server/constants"
	"log"
	"os"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func Init(ctx context.Context) {
	server := GetServer()
	err := PublicRoutes(ctx, server)
	if err != nil {
		panic("Failed to initialise router")
		return
	}

	server.Run()
}

func GetServer() *Server {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		requestID := uuid.New().String()

		c.Set(constants.RequestID, requestID)
		c.Writer.Header().Set(constants.RequestID, requestID)
		c.Next()
	})

	server := &Server{
		Engine: r,
		Port:   os.Getenv("PORT"),
	}

	return server
}

func (server *Server) Run() {
	log.Println("Port", server.Port)
	err := server.Engine.Run()
	if err != nil {
		panic("Failed to run router")
		return
	}
}
