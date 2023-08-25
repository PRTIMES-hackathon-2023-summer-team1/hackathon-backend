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

	tourRepository := repository.NewTourRepository(db)
	userRepository := repository.NewUserRepository(db)
	bookingRepository := repository.NewBookingRepository(db)
	jtiRepository := repository.NewJTIRepository(db) // todo: redisに変更
	tourController := controllers.NewTourController(tourRepository)
	userController := controllers.NewUserController(userRepository, jtiRepository)
	bookingController := controllers.NewBookingController(bookingRepository, tourRepository, userRepository)

	JWTAuthMiddleware := middleware.NewJWTAuthMiddleware(jtiRepository)

	tourGroup := r.Group("/tours")
	{
		//ツアー情報検索
		tourGroup.GET("/search", tourController.SearchTour)

		//ツアー情報閲覧系
		tourGroup.GET("", tourController.GetAllTours)
		tourGroup.GET("/:tour_id", tourController.GetTour)

		//ツアー情報操作系
		tourGroup.POST("", JWTAuthMiddleware.JWTAuthHandler(), tourController.CreateTour)
		tourGroup.PUT("", JWTAuthMiddleware.JWTAuthHandler(), tourController.EditTour)
	}

	bookingGroup := r.Group("/bookings")
	{
		// ツアー予約の投稿
		bookingGroup.POST("", JWTAuthMiddleware.JWTAuthHandler(), bookingController.PostBooking)
		// ツアー予約の取得
		bookingGroup.GET("", JWTAuthMiddleware.JWTAuthHandler(), bookingController.GetBookingByUserID)
		// ツアー予約の削除
		bookingGroup.DELETE("/:bookingID", JWTAuthMiddleware.JWTAuthHandler(), bookingController.DeleteBooking)
	}

	userGroup := r.Group("/users")
	{
		userGroup.POST("/signup", userController.Signup)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/refresh", userController.Refresh)
		userGroup.GET("/is_admin", JWTAuthMiddleware.JWTAuthHandler(), userController.IsAdmin)
	}

	return r
}
