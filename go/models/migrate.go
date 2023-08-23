package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&TestModel{})
}

func InsertDummyData(db *gorm.DB) {
	var test = []TestModel{
		{UID: "33u@2", Name: "Yuta"},
	}
	db.Save(&test)
}
