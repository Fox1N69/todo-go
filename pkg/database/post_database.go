package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostDB() *gorm.DB {
	dsn := "user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Error connect to database: ", err)
	}

	return DB
}
