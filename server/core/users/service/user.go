package service

import (
	"context"
	"github.com/projects/watch-list/server/core/commons"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	svcRequest "github.com/projects/watch-list/server/core/users/requests/service"
	commonError "github.com/projects/watch-list/server/pkg/error"
	"log"
	"sync"
)

type UserService struct {
	repo interfaces.UserRepository
}

var svc *UserService
var svcOnce sync.Once

func NewUserService(repo interfaces.UserRepository) *UserService {
	svcOnce.Do(func() {
		svc = &UserService{
			repo: repo,
		}
	})
	return svc
}

func (svc *UserService) GetUser(ctx context.Context,
	userID string,
) (res interface{}, cusErr commonError.CustomError) {
	logTag := commons.GetLogTag(ctx, "[UserService] GetUser")

	res, cusErr = svc.repo.GetUser(ctx, userID)
	if cusErr.Exists() {
		log.Printf(logTag+"unable to fetch user, err : %+v\n", cusErr.Error())
		return
	}

	return
}

func (svc *UserService) GetAllUsers(ctx context.Context) (res interface{}, cusErr commonError.CustomError) {
	logTag := commons.GetLogTag(ctx, "[UserService] GetAllUsers")

	res, cusErr = svc.repo.GetAllUsers(ctx, svcRequest.GetUserFilter())
	if cusErr.Exists() {
		log.Printf(logTag+"unable to fetch all users, err : %+v\n", cusErr.Error())
		return
	}

	return
}
