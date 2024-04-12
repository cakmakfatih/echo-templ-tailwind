package app

import (
	"gohtmx/internal/logging"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server interface {
	Run()
}

type server struct {
	logger *logging.Logger
	e      *echo.Echo
}

func (s *server) Run() {
	s.e.Start(":3000")
}

func staticCacheMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Request().URL.Path, "/assets") {
				c.Response().Header().Set("Cache-Control", "public, max-age=86399")
			}

			return next(c)
		}
	}
}

func NewServer(logger *logging.Logger) Server {
	s := &server{}

	s.logger = logger
	s.e = echo.New()

	s.e.Use(middleware.Gzip())
	s.e.Use(staticCacheMiddleware())

	s.e.Static("/static/", "./web/static/")

	routes(s)

	return s
}
