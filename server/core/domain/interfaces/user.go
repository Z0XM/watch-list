package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/projects/watch-list/server/core/user/service/requests"
)

type UserController interface {
	Login(ctx *gin.Context)
}

type UserService interface {
	Login(ctx context.Context, request requests.SvcLoginRequest) (response interface{}, err error)
}

type UserRepository interface {
	Login()
}
