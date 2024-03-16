package middleware

import (
	"github.com/elliot14A/siteforge/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v5"
	"github.com/spf13/viper"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("siteforge")
		if err != nil {
			return c.Redirect(302, "/dashboard/login")
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			secret := viper.GetString("jwt_secret")
			if secret == "" {
				secret = "secret"
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Redirect(302, "/dashboard/login")
		}

		return next(c)
	}
}
