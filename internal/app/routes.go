package app

import (
	"gohtmx/internal/handler"
	"gohtmx/internal/middleware"
	"gohtmx/internal/repository"
)

func routes(s *server) {
	(*s.logger).Info("Initializing routes")

	userRepository := repository.NewUserRepository(s.logger, s.db)

	registerIndexHandler(s)
	registerAuthHandler(s, &userRepository)
}

func registerIndexHandler(s *server) {
	(*s.logger).Info("Registering indexHandler to the route")
	indexHandler := handler.NewIndexHandler(s.logger)

	s.e.GET("/", indexHandler.Get, middleware.AuthGuard())
}

func registerAuthHandler(s *server, userRepository *repository.UserRepository) {
	(*s.logger).Info("Registering authHandler to the route")
	authHandler := handler.NewAuthHandler(s.logger, userRepository)

	s.e.GET("/login", authHandler.LoginPage, middleware.AuthGuard())
	s.e.POST("/auth/sign-in", authHandler.AuthenticateWithEmailAndPassword)
}
