package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/projects/watch-list/server/core/domain/models"
	commonError "github.com/projects/watch-list/server/pkg/error"
)

type UserController interface {
	GetMyUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}

type UserService interface {
	GetUser(ctx context.Context,
		userID string,
	) (interface{}, commonError.CustomError)

	GetAllUsers(ctx context.Context,
	) (interface{}, commonError.CustomError)
}

type UserRepository interface {
	GetUser(ctx context.Context,
		userID string,
	) (models.User, commonError.CustomError)

	GetAllUsers(ctx context.Context,
		filter map[string]interface{},
	) ([]models.User, commonError.CustomError)
}
