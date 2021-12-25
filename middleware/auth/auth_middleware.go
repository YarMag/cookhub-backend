package auth

import (
	"net/http"
	"context"
	"strings"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	HandleAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type FirebaseAuthMiddleware struct {
	Client *auth.Client
}

func (m FirebaseAuthMiddleware) HandleAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))

		if idToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token identity!")
		}

		token, err := m.Client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Can't check token")
		}
		c.Request().Header.Set("UUID", token.UID)
		return next(c)
	}
}