package main

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"github.com/projects/watch-list/server/database"
	"github.com/projects/watch-list/server/pkg/validator"
	"github.com/projects/watch-list/server/router"
)

const (
	modeHttp      = "http"
	modeMigration = "migration"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Not able to load env file")
	}

	ctx := context.Background()

	database.Init() // Initialising Database
	validator.Set() // Initialising Validator

	var mode string
	flag.StringVar(
		&mode,
		"mode",
		modeHttp,
		"Pass the flag to run in different modes (worker or migration)",
	)
	flag.Parse()

	switch mode {
	case modeHttp:
		router.Init(ctx) // Initialising Router
	case modeMigration:
		database.Migrate() // Running migrations
	default:
		router.Init(ctx) // Initialising Router
	}
}
