package api

import (
	"net/http"
	"regexp"

	"github.com/elliot14A/siteforge/actions/user"
	"github.com/elliot14A/siteforge/errors"
	"github.com/elliot14A/siteforge/pkg"
	"github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/labstack/echo/v5"
)

func login(ctx echo.Context) error {
	// get username and password from form

	email := ctx.FormValue("email")
	_ = ctx.FormValue("password")

	// check if email is a valid email
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	if !match {
		return pkg.Render(ctx, dashboard.LoginError("Enter a valid email"))
	}

	_, err := user.UserByEmail(email)

	if err != nil {
		if err == errors.ErrNotFound {
			return pkg.Render(ctx, dashboard.LoginError("Invalid email or password"))
		}
		return ctx.NoContent(500)
	}

	cookie := &http.Cookie{}
	cookie.Name = "siteforge"
	cookie.Value = ""
	ctx.SetCookie(cookie)
	ctx.Response().Header().Set("HX-Redirect", "/dashboard")

	return ctx.NoContent(302)
}
