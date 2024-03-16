package dashboard

import (
	"github.com/elliot14A/siteforge/pkg"
	dashboard_view "github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/labstack/echo/v5"
)

func aboutUsContent(ctx echo.Context) error {
	return pkg.Render(ctx, dashboard_view.AboutUs())
}
