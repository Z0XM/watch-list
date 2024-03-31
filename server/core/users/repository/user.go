package repository

import (
	"context"
	"github.com/projects/watch-list/server/constants"
	"github.com/projects/watch-list/server/core/commons"
	"github.com/projects/watch-list/server/core/domain/models"
	"github.com/projects/watch-list/server/database"
	commonError "github.com/projects/watch-list/server/pkg/error"
	"log"
	"sync"
)

type UserRepository struct {
	db *database.DbCluster
}

var repo *UserRepository
var repoOnce sync.Once

func NewUserRepository(db *database.DbCluster) *UserRepository {
	repoOnce.Do(func() {
		repo = &UserRepository{
			db: db,
		}
	})

	return repo
}

func (repo *UserRepository) GetUser(ctx context.Context, userID string) (res models.User, cusErr commonError.CustomError) {
	logTag := commons.GetLogTag(ctx, "[UserRepository] GetUser")

	result := repo.db.Master.Model(&models.User{}).Where("id = ? AND is_active = ?", userID, true).Find(&res)
	if resultErr := result.Error; resultErr != nil {
		log.Printf(logTag+"unable to fetch user, err : %+v\n", resultErr.Error())

		cusErr = commonError.NewCustomError(constants.BadRequest, resultErr.Error())
		return
	}

	return
}

func (repo *UserRepository) GetAllUsers(ctx context.Context, filter map[string]interface{}) (res []models.User, cusErr commonError.CustomError) {
	logTag := commons.GetLogTag(ctx, "[UserRepository] GetAllUsers")

	result := repo.db.Master.Model(&models.User{}).Scopes(commons.Paginate(ctx)).Where(filter).Find(&res)
	if resultErr := result.Error; resultErr != nil {
		log.Printf(logTag+"unable to fetch all users, err : %+v\n", resultErr.Error())

		cusErr = commonError.NewCustomError(constants.BadRequest, resultErr.Error())
		return
	}

	return
}
