package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbCluster struct {
	master *gorm.DB
}

type Db struct {
	Cluster *DbCluster
}

var dbInstance *Db

func Init() {
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		username,
		password,
		host,
		port,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	log.Println("Database connection successful")

	dbInstance = &Db{
		Cluster: &DbCluster{master: db},
	}
}

func GetCluster() *Db {
	return dbInstance
}
