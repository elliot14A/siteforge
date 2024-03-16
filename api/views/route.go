package views

import (
	"github.com/elliot14A/siteforge/api/views/dashboard"
	"github.com/labstack/echo/v5"
)

func InitViewRoutes(e *echo.Echo) {
	dashboard.InitDashboardRoutes(e)
}
