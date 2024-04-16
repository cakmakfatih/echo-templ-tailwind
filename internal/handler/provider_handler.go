package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
)

type ProviderHandler interface{}

type providerHandler struct {
	logger             *logging.Logger
	providerRepository *repository.ProviderRepository
}

func NewProviderHandler(logger *logging.Logger, providerRepository *repository.ProviderRepository) ProviderHandler {
	return &providerHandler{
		logger:             logger,
		providerRepository: providerRepository,
	}
}
