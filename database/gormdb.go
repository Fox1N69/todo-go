package database

import (
	"os"
	"path/filepath"
	"rest/models"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() *gorm.DB {
	if err := godotenv.Load(filepath.Join("../", ".env")); err != nil {
		logrus.Fatal("Gorm.io database", err)
	}

	DB, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Database connect...")

	if err := DB.AutoMigrate(&models.Users{}, &models.Posts{}); err != nil {
		logrus.Error("AutoMigrate not working")
	} else {
		logrus.Println("Database migrate!")
	}

	return DB
}
