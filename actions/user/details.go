package user

import (
	"fmt"

	"github.com/elliot14A/siteforge/errors"
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
)

func UserByEmail(email string) (*models.User, error) {
	db := pkg.GetDB()

	user := models.User{}

	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		fmt.Println("==>", err)
		return nil, errors.ErrNotFound
	}

	return &user, nil
}
