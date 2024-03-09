package cmd

import (
	"fmt"

	"github.com/elliot14A/siteforge/templates/dashboard"
	"github.com/elliot14A/siteforge/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts http server",
	Run: func(cmd *cobra.Command, args []string) {
		http()
	},
}

func http() {
	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
			return nil
		},
	}))
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, dashboard.Login())
	})
	e.Logger.Fatal(e.Start(":3000"))
}
