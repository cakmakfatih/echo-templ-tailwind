package app

import (
	"gohtmx/util"
	template "gohtmx/web/template/layout"
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes(s *server) {
	s.e.GET("/", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, template.MainLayout("Home"))
	})
}
