package router

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/controllers"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/middleware"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewRouter(db *gorm.DB) *gin.Engine {
	g := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	g.Use(cors.New(config))
	g.Use(middleware.ErrorHandler())

	testRepository := repository.NewTestRepository(db)
	testController := controllers.NewTestController(testRepository)
	g.POST("/test", testController.Set)
	return g
}
