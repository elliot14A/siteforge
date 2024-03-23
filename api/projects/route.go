package projects

import "github.com/labstack/echo/v5"

func InitViewRoutes(router *echo.Group) {
	routes := router.Group("/projects")
	routes.POST("/add", add)
}
