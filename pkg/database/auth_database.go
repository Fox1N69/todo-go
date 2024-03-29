package database

import (
	"blog/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitAuthDB() *gorm.DB {
	dsn := "user=postgres password=8008 dbname=sitesb port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.User{})

	return DB
}
