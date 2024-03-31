package service

import (
	"github.com/projects/watch-list/server/constants"
)

func GetUserFilter() map[string]interface{} {
	res := make(map[string]interface{})

	res[constants.IsActive] = true

	return res
}
