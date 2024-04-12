package handler

import (
	"gohtmx/util"
	template "gohtmx/web/template/layout"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexHandler interface {
	Get(echo.Context) error
}

type indexHandler struct{}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (*indexHandler) Get(c echo.Context) error {
	return util.Render(c, http.StatusOK, template.MainLayout("Home"))
}
