package dashboard

import (
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	dashboard_view "github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v5"
	"github.com/spf13/viper"
)

func loginView(ctx echo.Context) error {

	cookie, _ := ctx.Cookie("siteforge")

	if cookie != nil {
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			secret := viper.GetString("jwt_secret")
			if secret == "" {
				secret = "secret"
			}
			return []byte(secret), nil
		})

		if err == nil || token.Valid {
			return ctx.Redirect(302, "/dashboard/home")
		}
	}

	return pkg.Render(ctx, dashboard_view.Login())
}
