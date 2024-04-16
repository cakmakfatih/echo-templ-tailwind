package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"net/http"

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

func (h *panelHandler) Create(c echo.Context) error {
	var panelForm repository.PanelForm
	err := c.Bind(&panelForm)

	if err != nil {
		(*h.logger).Warn(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	user := c.Get("user").(*entity.UserSession)

	_, err = (*h.panelRepository).Create(user, &panelForm)

	if err != nil {
		(*h.logger).Warn(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
