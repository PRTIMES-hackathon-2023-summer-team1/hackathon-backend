package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MaxRetry = 50
	WaitTime = 1
)

func Connect(postgresInfo *config.PostgresInfo) (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", postgresInfo.Host, postgresInfo.User, postgresInfo.Password, postgresInfo.Database, postgresInfo.Port)
	var err error
	var db *gorm.DB
	for i := 0; i < MaxRetry; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("retrying to connect to db: %d", i)
		time.Sleep(WaitTime * time.Second)
	}
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	return db, sqlDB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.TestModel{})
}
