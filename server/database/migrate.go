package database

import (
	"github.com/projects/watch-list/server/core/domain/models"
	"log"
)

func Migrate() {
	log.Println("Running Migrations")

	err := GetCluster().Cluster.Master.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate user table")
	}
}
