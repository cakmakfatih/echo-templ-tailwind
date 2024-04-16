package handler

import (
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/repository"
	"gohtmx/util"
	page "gohtmx/web/template/page"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthHandler interface {
	AuthenticateWithEmailAndPassword(c echo.Context) error
	LoginPage(c echo.Context) error
}

type authHandler struct {
	logger          *logging.Logger
	userRepository  *repository.UserRepository
	panelRepository *repository.PanelRepository
}

func NewAuthHandler(logger *logging.Logger, userRepository *repository.UserRepository, panelRepository *repository.PanelRepository) AuthHandler {
	return &authHandler{
		logger:          logger,
		userRepository:  userRepository,
		panelRepository: panelRepository,
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

	panels, err := (*h.panelRepository).GetPanelsOfUser(user)

	if err != nil {
		(*h.logger).Warn("Error getting panels on [AuthenticateWithEmailAndPassword]")
	}

	var panelEntities []*entity.PanelSession

	for _, p := range panels {
		panelEntities = append(panelEntities, &entity.PanelSession{
			Id:              p.Id,
			User:            p.User,
			LoginURL:        p.LoginURL,
			SupportUsername: p.SupportUsername,
			SupportPassword: p.SupportPassword,
			TelegramToken:   p.TelegramToken,
			WhatsappToken:   p.WhatsappToken,
			Created:         p.Created.Time(),
			Updated:         p.Updated.Time(),
		})
	}

	sess.Values["panels"] = panelEntities

	err = sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Response().Header().Set("HX-Redirect", "/")

	return c.NoContent(http.StatusOK)
}

func (h *authHandler) LoginPage(c echo.Context) error {
	if c.Get("is_authenticated") == true {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}

	return util.Render(c, http.StatusOK, page.LoginPage())
}
