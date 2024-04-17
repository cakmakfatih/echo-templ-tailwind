package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceHandler interface {
	Get(c echo.Context) error
}

type serviceHandler struct {
	logger            *logging.Logger
	serviceRepository *repository.ServiceRepository
}

func NewServiceHandler(logger *logging.Logger, serviceRepository *repository.ServiceRepository) ServiceHandler {
	return &serviceHandler{
		logger:            logger,
		serviceRepository: serviceRepository,
	}
}

func (h *serviceHandler) Get(c echo.Context) error {
	panels := c.Get("panels").([]*entity.PanelSession)
	_, err := (*h.serviceRepository).Get(panels)

	if err != nil {
		(*h.logger).Warn("Error occurred on serviceHandler/get")
		(*h.logger).Warn(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
