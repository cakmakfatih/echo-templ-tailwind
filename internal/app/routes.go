package app

import (
	"gohtmx/internal/handler"
	"gohtmx/internal/middleware"
	"gohtmx/internal/repository"
)

func routes(s *server) {
	(*s.logger).Info("Initializing routes")

	userRepository := repository.NewUserRepository(s.logger, s.db)
	panelRepository := repository.NewPanelRepository(s.logger, s.db)
	serviceRepository := repository.NewServiceRepository(s.logger, s.db)
	providerRepository := repository.NewProviderRepository(s.logger, s.db)

	registerIndexHandler(s)
	registerAuthHandler(s, &userRepository, &panelRepository)
	registerPanelHandler(s, &panelRepository)
	registerServiceHandler(s, &serviceRepository)
	registerPartialHandler(s, &providerRepository)
}

func registerIndexHandler(s *server) {
	(*s.logger).Info("Registering indexHandler to the route")
	indexHandler := handler.NewIndexHandler(s.logger)

	s.e.GET("/", indexHandler.Get, middleware.AuthGuard())
}

func registerAuthHandler(s *server, userRepository *repository.UserRepository, panelRepository *repository.PanelRepository) {
	(*s.logger).Info("Registering authHandler to the route")
	authHandler := handler.NewAuthHandler(s.logger, userRepository, panelRepository)

	s.e.GET("/login", authHandler.LoginPage, middleware.AuthGuard())
	s.e.POST("/auth/sign-in", authHandler.AuthenticateWithEmailAndPassword)
	s.e.GET("/auth/sign-out", authHandler.Logout)
}

func registerPanelHandler(s *server, panelRepository *repository.PanelRepository) {
	(*s.logger).Info("Registering panelHandler to the route")
	panelHandler := handler.NewPanelHandler(s.logger, panelRepository)

	s.e.POST("/panel", panelHandler.Create, middleware.AuthGuard())
}

func registerServiceHandler(s *server, serviceRepository *repository.ServiceRepository) {
	(*s.logger).Info("Registering serviceHandler to the route")
	serviceHandler := handler.NewServiceHandler(s.logger, serviceRepository)

	s.e.GET("/service", serviceHandler.Get, middleware.AuthGuard())
}

func registerPartialHandler(s *server, providerRepository *repository.ProviderRepository) {
	(*s.logger).Info("Registering partialHandler to the route")
	partialHandler := handler.NewPartialHandler(s.logger, providerRepository)

	s.e.GET("/partial/fragment/providers", partialHandler.GetProvidersFragment, middleware.AuthGuard())
}
