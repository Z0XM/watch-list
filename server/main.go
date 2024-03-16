package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/projects/watch-list/server/database"
	"github.com/projects/watch-list/server/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Not able to load config file")
	}

	ctx := context.Background()

	// Initialising Database
	database.Init()

	// Initialising Router
	router.Init(ctx)
}
