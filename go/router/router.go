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
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.Use(middleware.ErrorHandler())

	testRepository := repository.NewTestRepository(db)
	testController := controllers.NewTestController(testRepository)
	r.POST("/test", testController.Set)

	tourGroup := r.Group("/tours")
	{
		tourRepository := repository.NewTourRepository(db)
		tourController := controllers.NewTourController(tourRepository)
		tourGroup.GET("", tourController.GetAll)
		tourGroup.GET("/:tour_id", tourController.Get)
		tourGroup.POST("", tourController.CreateTour)
	}

	return r
}
