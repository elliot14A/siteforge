package cmd

import (
	"fmt"
	"log"

	"github.com/elliot14A/siteforge/api"
	"github.com/elliot14A/siteforge/pkg"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts http server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	app := echo.New()

	pkg.SetupLogger()
	pkg.SetupDB()

	app.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				fmt.Printf("REQUEST: uri: %v, status: %v, latency: %v\n", v.URI, v.Status, v.Latency.Milliseconds())
				return nil
			},
		}),
	)
	app.Static("/static", "static")
	api.InitRoutes(app)
	log.Fatal(app.Start(":3000"))
}
