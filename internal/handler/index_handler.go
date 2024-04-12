package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/util"
	template "gohtmx/web/template/layout"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexHandler interface {
	Get(echo.Context) error
}

type indexHandler struct {
	logger *logging.Logger
}

func NewIndexHandler(logger *logging.Logger) IndexHandler {
	return &indexHandler{
		logger: logger,
	}
}

func (*indexHandler) Get(c echo.Context) error {
	return util.Render(c, http.StatusOK, template.MainLayout("Home"))
}
