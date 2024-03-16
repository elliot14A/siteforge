package cmd

import (
	"fmt"

	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(adminCmd)
}

var adminCmd = &cobra.Command{
	Use:   "create-admin",
	Short: "Create admin user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: create-admin <email> <password>")
			return
		}
		create_admin(args[0], args[1])
	},
}

func create_admin(email string, password string) {
	pkg.SetupLogger()
	pkg.SetupDB()

	// check if admin is already created

	db := pkg.GetDB()

	var count int64

	err := db.Model(&models.User{}).Count(&count).Error

	if err != nil {
		pkg.GetLogger().Error("Error creating admin")
		return
	}

	if count > 0 {
		pkg.GetLogger().Info("Admin already created")
		return
	}

	hashedPassword := pkg.HashPassword(password)

	admin := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = db.Create(&admin).Error

	if err != nil {
		pkg.GetLogger().Error("Error creating admin")
		return
	}
}
