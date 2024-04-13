package main

import (
	"gohtmx/config"
	"gohtmx/internal/app"
	"gohtmx/internal/logging"
)

func main() {
	config.InitConfig()
	logger := logging.NewLogger()
	server := app.NewServer(&logger)

	server.Run()
}
