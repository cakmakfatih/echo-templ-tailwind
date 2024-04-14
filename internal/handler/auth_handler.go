package handler

import (
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"gohtmx/util"
	page "gohtmx/web/template/page"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	GetLogin(c echo.Context) error
	AuthenticateWithEmailAndPassword(c echo.Context) error
}

type authHandler struct {
	logger         *logging.Logger
	userRepository *repository.UserRepository
}

func NewAuthHandler(logger *logging.Logger, userRepository *repository.UserRepository) AuthHandler {
	return &authHandler{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (*authHandler) GetLogin(c echo.Context) error {
	return util.Render(c, http.StatusOK, page.LoginPage())
}

func (h *authHandler) AuthenticateWithEmailAndPassword(c echo.Context) error {
	var loginForm repository.LoginForm
	err := c.Bind(&loginForm)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if loginForm.Email == "" || loginForm.Password == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	(*h.userRepository).AuthenticateWithEmailAndPassword(&loginForm)

	return c.NoContent(http.StatusOK)
}
