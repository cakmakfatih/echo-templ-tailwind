package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"gohtmx/util"
	"net/http"

	fragment "gohtmx/web/template/partial/fragment"

	"github.com/labstack/echo/v4"
)

type PartialHandler interface {
	GetProvidersFragment(c echo.Context) error
	GetServicesOfProvider(c echo.Context) error
}

type partialHandler struct {
	logger             *logging.Logger
	providerRepository *repository.ProviderRepository
	serviceRepository  *repository.ServiceRepository
}

func NewPartialHandler(logger *logging.Logger, providerRepository *repository.ProviderRepository, serviceRepository *repository.ServiceRepository) PartialHandler {
	return &partialHandler{
		logger:             logger,
		providerRepository: providerRepository,
		serviceRepository:  serviceRepository,
	}
}

func (h *partialHandler) GetProvidersFragment(c echo.Context) error {
	panels := c.Get("panels").([]*entity.PanelSession)
	providers, err := (*h.providerRepository).Get(panels)

	if err != nil {
		(*h.logger).Warn(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	return util.Render(c, http.StatusOK, fragment.Providers(panels, providers))
}

func (h *partialHandler) GetServicesOfProvider(c echo.Context) error {
	providerId := c.Param("providerId")

	if providerId == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err := (*h.serviceRepository).GetFromProvider(providerId)

	if err != nil {
		(*h.logger).Warn(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
