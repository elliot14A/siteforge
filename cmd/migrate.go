package cmd

import (
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	pkg.SetupLogger()
	logger := pkg.GetLogger()
	logger.Info("Running migrations")
	pkg.SetupDB()
	db := pkg.GetDB()

	err := db.AutoMigrate(&models.User{})

	if err != nil {
		logger.Error("Error running migrations")
		return
	}

	logger.Info("Migrations complete")
}
