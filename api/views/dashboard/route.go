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
	content := routes.Group("/content")
	content.GET("/home", homeContent)
	content.GET("/projects", projectsContent)
	content.GET("/certifications", certificationsContent)
	content.GET("/gallery", galleryContent)
	content.GET("/aboutus", aboutUsContent)
	content.GET("/awards", awardsContent)
}
