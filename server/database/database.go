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
	*DbCluster
}

var dbInstance *Db

func Init() {
	host := os.Getenv("HOST")
	//port := os.Getenv("DB_PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require",
		username,
		password,
		host,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	log.Println("Database connection successful")

	SetCluster(&DbCluster{master: db})
}

func GetCluster() *Db {
	return dbInstance
}

func SetCluster(cluster *DbCluster) {
	dbInstance = &Db{cluster}
}
