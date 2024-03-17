package database

import "github.com/projects/watch-list/server/core/domain/models"

func Migrate() {
	err := GetCluster().Cluster.master.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate user table")
	}
}
