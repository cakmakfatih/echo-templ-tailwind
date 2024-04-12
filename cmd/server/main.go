package main

import (
	"gohtmx/internal/app"
	"gohtmx/internal/logging"
)

func main() {
	logger := logging.NewLogger()
	server := app.NewServer(&logger)

	server.Run()
}
