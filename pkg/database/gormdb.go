package database

import (
	"rest/pkg/models"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "user=postgres password=8008 dbname=sitesb port=5432 sslmode=disable"

func InitGormDB() *gorm.DB {
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database", err)
	}

	DB.AutoMigrate(
		&models.Users{},
		&models.Posts{},
	)

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		sleep := time.Second
		for DB == nil {
			DB = InitGormDB()
			time.Sleep(sleep)
			sleep *= 2
		}
	}

	return DB
}

// create func for update database
func UpdateDB() *gorm.DB {

	return nil
}
