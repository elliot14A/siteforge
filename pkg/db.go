package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func SetupDB() error {
	logger := GetLogger()
	logger.Info("Initializing database")
	DB, _ = gorm.Open(sqlite.Open("siteforge.db"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	logger.Info("Database initialized")
	return nil
}
