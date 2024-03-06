package database

import (
	"os"
	"rest/models"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN = "user=postgres password=8008 dbname=sitesb port=5432 sslmode=disable"

func InitGormDB() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	DB, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connect...")

	if err := DB.AutoMigrate(&models.Users{}, &models.Posts{}); err != nil {
		log.Error("AutoMigrate not working")
	} else {
		log.Println("Database migrate!")
	}

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitGormDB()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			log.Println("Database is unavaibl. Wait for %d sec.\n")
			time.Sleep(sleep * time.Second)
			DB = InitGormDB()
		}
	}

	return DB
}
