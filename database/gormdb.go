package database

import (
	"os"
	"path/filepath"
	"rest/models"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() *gorm.DB {
	if err := godotenv.Load(filepath.Join(".env")); err != nil {
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

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitGormDB()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			logrus.Println("Database is unavaibl. Wait for %d sec.\n")
			time.Sleep(sleep * time.Second)
			DB = InitGormDB()
		}
	}

	return DB
}
