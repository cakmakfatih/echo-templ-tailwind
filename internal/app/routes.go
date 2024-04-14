package app

import (
	"gohtmx/internal/handler"
	"net/http"
)

func routes(s *server) {
	(*s.logger).Info("Initializing routes")

	registerIndexHandler(s)
	registerAuthHandler(s)
}

func registerIndexHandler(s *server) {
	(*s.logger).Info("Registering indexHandler to the route")
	indexHandler := handler.NewIndexHandler(s.logger)

	s.e.Add(http.MethodGet, "/", indexHandler.Get)
}

func registerAuthHandler(s *server) {
	(*s.logger).Info("Registering authHandler to the route")
	authHandler := handler.NewAuthHandler(s.logger)

	s.e.Add(http.MethodGet, "/login", authHandler.GetLogin)
}
