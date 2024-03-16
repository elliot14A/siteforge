package dashboard

import (
	"github.com/elliot14A/siteforge/api/middleware"
	"github.com/labstack/echo/v5"
)

func InitDashboardRoutes(router *echo.Echo) {
	routes := router.Group("/dashboard")
	routes.GET("/login", loginView)
	routes.Use(middleware.Authorize)
	routes.GET("/home", homeView)
	routes.GET("/", func(_ echo.Context) error { return nil })
}
