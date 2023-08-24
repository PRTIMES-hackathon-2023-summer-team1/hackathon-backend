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
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))
	r.Use(middleware.ErrorHandler())

	testRepository := repository.NewTestRepository(db)
	tourRepository := repository.NewTourRepository(db)
	userRepository := repository.NewUserRepository(db)
	bookingRepository := repository.NewBookingRepository(db)
	testController := controllers.NewTestController(testRepository)
	tourController := controllers.NewTourController(tourRepository)
	userController := controllers.NewUserController(userRepository)
	bookingController := controllers.NewBookingController(bookingRepository, tourRepository, userRepository)

	r.POST("/test", testController.Set)
	tourGroup := r.Group("/tours")
	{
		//ツアー情報閲覧系
		tourGroup.GET("", tourController.GetAllTours)
		tourGroup.GET("/:tour_id", tourController.GetTour)

		//ツアー情報操作系
		tourGroup.POST("", tourController.CreateTour)
		tourGroup.PUT("", tourController.EditTour)
	}

	bookingGroup := r.Group("/booking")
	{
		// ツアー予約の投稿
		bookingGroup.POST("/", bookingController.PostBooking)
		// ツアー予約の取得
		bookingGroup.GET("/:userID", bookingController.GetBookingByUserID)
		// ツアー予約の削除
		bookingGroup.DELETE("/:bookingID", bookingController.DeleteBooking)
	}

	userGroup := r.Group("/users")
	{
		userGroup.POST("/signup", userController.Signup)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/is_admin", userController.IsAdmin)
	}

	return r
}
