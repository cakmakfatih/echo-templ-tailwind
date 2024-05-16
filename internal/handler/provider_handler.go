package handler

import (
	"github.com/labstack/echo/v4"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"net/http"
)

type ProviderHandler interface {
	Create(c echo.Context) error
}

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

func (h *providerHandler) Create(c echo.Context) error {
	var providerForm repository.ProviderForm
	err := c.Bind(&providerForm)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = (*h.providerRepository).Create(&providerForm)

	if err != nil {
		(*h.logger).Warn(err.Error())
	}

	return c.NoContent(http.StatusOK)
}
