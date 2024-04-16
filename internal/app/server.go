package app

import (
	"encoding/gob"
	"fmt"
	"gohtmx/internal/database"
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"log/slog"
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
	db     *database.DB
}

func (s *server) Run() {
	(*s.logger).Info("Starting server on port", slog.String("port", os.Getenv("PORT")))
	s.e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

func staticCacheMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Request().URL.Path, "/static") {
				c.Response().Header().Set("Cache-Control", "public, max-age=86399")
			}

			return next(c)
		}
	}
}

func NewServer(logger *logging.Logger, db *database.DB) Server {
	(*logger).Info("Initializing Server")

	gob.Register(&entity.UserSession{})
	gob.Register([]*entity.PanelSession{})

	s := &server{
		logger: logger,
		e:      echo.New(),
		db:     db,
	}

	(*logger).Info("Assigning middlewares")

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
