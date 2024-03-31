//go:build wireinject
// +build wireinject

package users

import (
	"github.com/google/wire"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	"github.com/projects/watch-list/server/database"
)

func Wire(db *database.DbCluster) (interfaces.UserController, error) {
	panic(wire.Build(ProviderSet))
}
