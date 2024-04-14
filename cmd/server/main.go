package main

import (
	"gohtmx/config"
	"gohtmx/internal/app"
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
)

func main() {
	config.InitConfig()
	db := database.NewDB()
	logger := logging.NewLogger()
	server := app.NewServer(&logger, db)

	server.Run()
}
