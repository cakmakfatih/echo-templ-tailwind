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
}

type partialHandler struct {
	logger             *logging.Logger
	providerRepository *repository.ProviderRepository
}

func NewPartialHandler(logger *logging.Logger, providerRepository *repository.ProviderRepository) PartialHandler {
	return &partialHandler{
		logger:             logger,
		providerRepository: providerRepository,
	}
}

func (h *partialHandler) GetProvidersFragment(c echo.Context) error {
	panels := c.Get("panels").([]*entity.PanelSession)
	providers, err := (*h.providerRepository).GetProviders(panels)

	if err != nil {
		return nil
	}

	return util.Render(c, http.StatusOK, fragment.Providers(panels, providers))
}
