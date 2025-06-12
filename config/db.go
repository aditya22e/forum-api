package config

import (
	"forum-api/models"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate models
	err = db.AutoMigrate(&models.User{}, &models.Topic{}, &models.Comment{}, &models.Subscription{})
	if err != nil {
		return err
	}

	DB = db
	logrus.Info("Database connected successfully")
	return nil
}
