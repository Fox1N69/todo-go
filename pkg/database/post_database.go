package database

import (
	"blog/pkg/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostDB() *gorm.DB {
	dsn := "user=postgres password=8008 dbname=sitesb port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Error connect to database: ", err)
	}

	DB.AutoMigrate(&models.Post{})

	return DB
}
