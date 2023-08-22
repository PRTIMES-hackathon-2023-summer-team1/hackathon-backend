package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(postgresInfo *config.PostgresInfo) (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", postgresInfo.Host, postgresInfo.User, postgresInfo.Password, postgresInfo.Database, postgresInfo.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	return db, sqlDB
}

func Migrate(db *gorm.DB){
	db.AutoMigrate(&models.TestModel{})
}