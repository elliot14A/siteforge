package api

import (
	"github.com/elliot14A/siteforge/api/projects"
	"github.com/elliot14A/siteforge/api/views"
	"github.com/labstack/echo/v5"
)

func InitRoutes(e *echo.Echo) {
	views.InitViewRoutes(e)
	api := e.Group("/api")
	api.POST("/login", login)
	projects.InitViewRoutes(api)
}
