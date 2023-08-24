package main

import (
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/router"
)

func main() {
	appInfo := config.LoadConfig()
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	repo, sqlDB := db.Connect(appInfo.PostgresInfo)
	defer sqlDB.Close()
	models.Drop(repo)
	models.Migrate(repo)
	models.InsertDummyData(repo)

	r := router.NewRouter(repo)
	r.Run(":8080")
}
