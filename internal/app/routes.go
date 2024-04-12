package app

import (
	"gohtmx/internal/handler"
)

func routes(s *server) {
	registerIndexHandlers(s)
}

func registerIndexHandlers(s *server) {
	indexHandler := handler.NewIndexHandler(s.logger)

	s.e.GET("/", indexHandler.Get)
}
