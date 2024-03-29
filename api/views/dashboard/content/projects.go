package content

import (
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	content_templ "github.com/elliot14A/siteforge/templates/dashboard/content"
	"github.com/labstack/echo/v5"
)

func projectsContent(ctx echo.Context) error {
	return pkg.Render(ctx, content_templ.Projects([]models.Project{}))
}

func add_projects(ctx echo.Context) error {
	return pkg.Render(ctx, content_templ.AddProject())
}
