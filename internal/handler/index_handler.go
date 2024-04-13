package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/util"
	page "gohtmx/web/template/page"
	"net/http"

	"github.com/labstack/echo/v5"
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
	return util.Render(c, http.StatusOK, page.HomePage())
}
