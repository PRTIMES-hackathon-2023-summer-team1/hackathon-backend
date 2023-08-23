package main

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/config"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/controllers"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/router"
)

func main() {
	appInfo := config.LoadConfig()
	repo, sqlDB := db.Connect(appInfo.PostgresInfo)
	defer sqlDB.Close()
	models.Migrate(repo)
	models.InsertDummyData(repo)

	r := router.NewRouter(repo)

	userRepo := repository.NewUserRepository(repo)
	userController := controllers.NewUserController(userRepo)
	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)

	r.Run(":8080")
}
