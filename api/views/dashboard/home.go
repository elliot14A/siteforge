package dashboard

import (
	"github.com/elliot14A/siteforge/pkg"
	dashboard_view "github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/labstack/echo/v5"
)

func homeView(ctx echo.Context) error {
	return pkg.Render(ctx, dashboard_view.Home())
}

func homeContent(ctx echo.Context) error {
	return pkg.Render(ctx, dashboard_view.HomeContent())
}
