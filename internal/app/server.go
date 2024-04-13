package app

import (
	"gohtmx/internal/logging"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	s.e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteLaxMode,
	}))
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET_TOKEN")))))

	s.e.Static("/static/", "./web/static/")

	routes(s)

	return s
}
