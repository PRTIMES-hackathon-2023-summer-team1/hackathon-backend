package main

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"
)

func main() {
	appInfo := config.LoadConfig()
	database, sqlDB := db.Connect(appInfo.PostgresInfo)
	defer sqlDB.Close()
	db.Migrate(database)
}
