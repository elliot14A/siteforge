package projects

import (
	"fmt"

	"github.com/elliot14A/siteforge/pkg"
	"github.com/elliot14A/siteforge/templates/dashboard/content"
	"github.com/labstack/echo/v5"
)

func add(ctx echo.Context) error {
	image, err := ctx.FormFile("image")
	if err != nil {
		fmt.Println(err)
	}
	location := ctx.FormValue("location")
	fmt.Println("==> ", image.Filename, location)
	return pkg.Render(ctx, content.StopLoading())
}
