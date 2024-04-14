package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/util"
	page "gohtmx/web/template/page"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	GetLogin(c echo.Context) error
}

type authHandler struct {
	logger *logging.Logger
}

func NewAuthHandler(logger *logging.Logger) AuthHandler {
	return &authHandler{
		logger: logger,
	}
}

func (*authHandler) GetLogin(c echo.Context) error {
	return util.Render(c, http.StatusOK, page.LoginPage())
}
