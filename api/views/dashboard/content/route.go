package content

import "github.com/labstack/echo/v5"

func InitDashboardContentRoutes(routes *echo.Group) {
	content := routes.Group("/content")
	content.GET("/home", homeContent)
	content.GET("/projects", projectsContent)
	content.GET("/certifications", certificationsContent)
	content.GET("/gallery", galleryContent)
	content.GET("/aboutus", aboutUsContent)
	content.GET("/awards", awardsContent)
}
