package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"

	"github.com/labstack/echo/v4"
)

type PanelHandler interface {
	Create(c echo.Context) error
}

type panelHandler struct {
	logger          *logging.Logger
	panelRepository *repository.PanelRepository
}

func NewPanelHandler(logger *logging.Logger, panelRepository *repository.PanelRepository) PanelHandler {
	return &panelHandler{
		logger:          logger,
		panelRepository: panelRepository,
	}
}

func (*panelHandler) Create(c echo.Context) error {
	return nil
}
