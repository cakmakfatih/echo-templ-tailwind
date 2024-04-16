package middleware

import (
	"gohtmx/internal/entity"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func redirectOrNextBasedOnPath(c echo.Context, path string, next echo.HandlerFunc) error {
	c.Set("is_authenticated", false)

	if path != "/login" {
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	} else {
		return next(c)
	}
}

func AuthGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqPath := c.Request().URL.Path
			sess, err := session.Get("session", c)

			if err != nil {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			user, exists := sess.Values["user"].(*entity.UserSession)

			if !exists {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			panels, exists := sess.Values["panels"].([]*entity.PanelSession)

			if !exists {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			savedCsrf, exists := sess.Values["csrf"].(string)

			if !exists {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			csrf, err := c.Cookie("_csrf")

			if err != nil {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			if savedCsrf != csrf.Value {
				return redirectOrNextBasedOnPath(c, reqPath, next)
			}

			c.Set("is_authenticated", true)
			c.Set("user", &user)
			c.Set("panels", panels)

			return next(c)
		}
	}
}
