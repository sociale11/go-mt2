package cmd

import (
	"go-mt2/config"
	"go-mt2/pkg/auth"
	"go-mt2/pkg/db"
)

type App struct {
	Logger Logger
	Auth   auth.AuthService
	DB     db.DBService
	Config config.Config
}

type Logger struct {
	Level string
}

func Execute() *App {

	dbService := db.DBService{
		AuthDB: db.NewAuthDB(),
		GameDB: db.NewGameDB(),
	}

	dbService.MigrateAuth()
	dbService.SeedAuth()

	dbService.MigrateGame()

	app := App{
		Logger: Logger{},
		Auth:   auth.AuthService{},
		DB:     db.DBService{},
		Config: config.Settings,
	}

	return &app
}
