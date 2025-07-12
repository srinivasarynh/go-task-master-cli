package database

import (
	"taskmaster/internal/config"
	"taskmaster/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Initialize() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.AppConfig.Database.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Task{}); err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
