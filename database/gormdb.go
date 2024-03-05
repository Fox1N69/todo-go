package database

import (
	"rest/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() *gorm.DB {
	dsn := "user=postgres password=8008 dbname=sitesb port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Database connect...")

	if err := DB.AutoMigrate(&models.Users{}); err != nil {
		logrus.Error("AutoMigrate not working")
	} else {
		logrus.Println("Database migrate!")
	}
	


	return DB
}