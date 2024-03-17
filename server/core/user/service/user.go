package service

import (
	"context"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	"github.com/projects/watch-list/server/core/user/service/requests"
	"sync"
)

type Service struct {
	repository interfaces.UserRepository
}

var svc *Service
var svcOnce sync.Once

func NewService(repo interfaces.UserRepository) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			repository: repo,
		}
	})
	return svc
}

func (us *Service) Login(ctx context.Context, request requests.SvcLoginRequest) (response interface{}, err error) {

}
