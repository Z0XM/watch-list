package users

import (
	"github.com/google/wire"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	"github.com/projects/watch-list/server/core/users/controller"
	"github.com/projects/watch-list/server/core/users/repository"
	"github.com/projects/watch-list/server/core/users/service"
)

var ProviderSet = wire.NewSet(
	controller.NewUserController,
	service.NewUserService,
	repository.NewUserRepository,

	// bind each one of the interfaces
	wire.Bind(new(interfaces.UserController), new(*controller.UserController)),
	wire.Bind(new(interfaces.UserService), new(*service.UserService)),
	wire.Bind(new(interfaces.UserRepository), new(*repository.UserRepository)),
)
