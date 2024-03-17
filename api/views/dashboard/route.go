package dashboard

import (
	"github.com/elliot14A/siteforge/api/middleware"
	"github.com/elliot14A/siteforge/api/views/dashboard/content"
	"github.com/labstack/echo/v5"
)

func InitDashboardRoutes(router *echo.Echo) {
	routes := router.Group("/dashboard")
	routes.GET("/login", loginView)
	routes.Use(middleware.Authorize)
	routes.GET("/home", homeView)
	content.InitDashboardContentRoutes(routes)
}
