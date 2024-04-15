package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthHandler interface {
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

func (h *authHandler) AuthenticateWithEmailAndPassword(c echo.Context) error {
	var loginForm repository.LoginForm
	err := c.Bind(&loginForm)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if loginForm.Email == "" || loginForm.Password == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := (*h.userRepository).AuthenticateWithEmailAndPassword(&loginForm)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	sess, err := session.Get("session", c)

	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user"] = &entity.UserSession{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Roles:    user.Roles,
	}
	sess.Values["csrf"] = c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)

	err = sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, user.Id)
}
