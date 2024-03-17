package content

import (
	"github.com/elliot14A/siteforge/pkg"
	content_templ "github.com/elliot14A/siteforge/templates/dashboard/content"
	"github.com/labstack/echo/v5"
)

func certificationsContent(ctx echo.Context) error {
	return pkg.Render(ctx, content_templ.Certificates())
}
