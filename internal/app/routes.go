package app

import (
	"gohtmx/internal/handler"
)

func routes(s *server) {
	(*s.logger).Info("Initializing routes")
	registerIndexHandlers(s)
}

func registerIndexHandlers(s *server) {
	(*s.logger).Info("Registering indexHandler to the route")
	indexHandler := handler.NewIndexHandler(s.logger)

	s.e.GET("/", indexHandler.Get)
}
