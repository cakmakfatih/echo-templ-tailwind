package app

import (
	"gohtmx/internal/logging"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Run()
}

type server struct {
	logger *logging.Logger
	e      *echo.Echo
}

func (s *server) Run() {
	s.e.Logger.Fatal(s.e.Start(":3000"))
}

func NewServer(logger *logging.Logger) Server {
	s := &server{}

	s.logger = logger
	s.e = echo.New()

	s.e.Static("/static/", "./web/static/")

	routes(s)

	return s
}
