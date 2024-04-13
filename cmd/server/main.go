package main

import (
	"gohtmx/config"
	"gohtmx/internal/app"
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
)

func main() {
	database.NewDB()
	config.InitConfig()
	logger := logging.NewLogger()
	server := app.NewServer(&logger)

	server.Run()
}
