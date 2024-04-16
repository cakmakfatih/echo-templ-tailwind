package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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

	panel, err := (*h.panelRepository).Create(user, &panelForm)

	if err != nil {
		(*h.logger).Warn(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	panels := c.Get("panels").([]*entity.PanelSession)
	panelEntity := &entity.PanelSession{
		Id:              panel.Id,
		User:            panel.User,
		LoginURL:        panel.LoginURL,
		SupportUsername: panel.SupportUsername,
		SupportPassword: panel.SupportPassword,
		TelegramToken:   panel.TelegramToken,
		WhatsappToken:   panel.WhatsappToken,
		Created:         panel.Created.Time(),
		Updated:         panel.Updated.Time(),
	}

	panels = append(panels, panelEntity)

	sess, err := session.Get("session", c)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["panels"] = panels

	err = sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Response().Header().Set("HX-Refresh", "true")

	return c.NoContent(http.StatusOK)
}
