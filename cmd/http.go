package cmd

import (
	"log"

	"github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/elliot14A/siteforge/utils"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts http server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDev: false,
	})
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/static", "static")
		e.Router.GET("/", func(c echo.Context) error {
			return utils.Render(c, dashboard.Login())
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
