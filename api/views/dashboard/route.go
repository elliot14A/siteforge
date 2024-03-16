package dashboard

import "github.com/labstack/echo/v5"

func InitDashboardRoutes(router *echo.Echo) {
	routes := router.Group("/dashboard")
	routes.GET("/login", dashboardView)
}
