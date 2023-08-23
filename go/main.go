package main

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/router"
)

func main() {
	appInfo := config.LoadConfig()
	_, sqlDB := db.Connect(appInfo.PostgresInfo)
	defer sqlDB.Close()
	models.Migrate()
	models.InsertDummyData()

	r := router.NewRouter()
	r.Run(":8080")
}
