package models

import (
	"gorm.io/gorm"
	"time"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&TestModel{}, &User{}, &Booking{}, &Tour{})
}

func InsertDummyData(db *gorm.DB) {
	var test = []TestModel{
		{UID: "33u@2", Name: "Yuta"},
	}

	var tour = []Tour{
		{
			TourID:          "tour123",
			UserID:          "user456",
			Name:            "Sample Tour",
			Description:     "This is a sample tour.",
			Body:            "Sample tour details.",
			Price:           100,
			FirstDay:        time.Now(),
			LastDay:         time.Now().AddDate(0, 0, 7),
			MaxCapacity:     20,
			CurrentCapacity: 10,
			IsVisible:       true,
		},
	}
	db.Save(&test)
	db.Save(&tour)
}
