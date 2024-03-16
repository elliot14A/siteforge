package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/elliot14A/siteforge/actions/user"
	"github.com/elliot14A/siteforge/errors"
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	"github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v5"
	"github.com/spf13/viper"
)

func login(ctx echo.Context) error {
	// get username and password from form

	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	// check if email is a valid email
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	if !match {
		return pkg.Render(ctx, dashboard.LoginError("Enter a valid email"))
	}

	user, err := user.UserByEmail(email)

	if err != nil {
		if err == errors.ErrNotFound {
			return pkg.Render(ctx, dashboard.LoginError("Invalid email or password"))
		}
		return ctx.NoContent(500)
	}

	err = pkg.VerifyPassword(user.Password, password)
	if err != nil {
		return pkg.Render(ctx, dashboard.LoginError("Invalid email or password"))
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{})

	secret := viper.GetString("jwt_secret")
	if secret == "" {
		secret = "secret"
	}

	token, err := tokenJwt.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err)
		return ctx.NoContent(500)
	}

	cookie := &http.Cookie{}
	cookie.Name = "siteforge"
	cookie.Value = token
	cookie.Path = "/"
	ctx.SetCookie(cookie)
	ctx.Response().Header().Set("HX-Redirect", "/dashboard/home")

	return ctx.NoContent(302)
}
