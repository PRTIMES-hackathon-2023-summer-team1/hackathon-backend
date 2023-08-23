package models

import "github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"

func Migrate() {
	db.DB.AutoMigrate(&Test{})
}

func InsertDummyData() {
	var test = []Test{
		{UID: "33u@2", Name: "Yuta"},
	}
	db.DB.Save(&test)
}
