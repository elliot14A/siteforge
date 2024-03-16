package user

import (
	"github.com/elliot14A/siteforge/errors"
	"github.com/elliot14A/siteforge/models"
	"github.com/elliot14A/siteforge/pkg"
	"gorm.io/gorm"
)

func UserByEmail(email string) (*models.User, error) {
	db := pkg.GetDB()

	user := models.User{}

	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrInvalidPassword
		}
		return nil, errors.ErrInternalError
	}

	return &user, nil
}
