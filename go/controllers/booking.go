package controllers

import (
	"log"
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingController struct {
	bookingModelRepository repository.IBookingRepository
	tourModelRepository    repository.ITourRepository
	userModelRepository    repository.IUserRepository
}

func NewBookingController(bRepo repository.IBookingRepository, tRepo repository.ITourRepository, uRepo repository.IUserRepository) *BookingController {
	return &BookingController{bookingModelRepository: bRepo, tourModelRepository: tRepo, userModelRepository: uRepo}
}

func (b BookingController) PostBooking(c *gin.Context) {
	var booking models.Booking
	err := c.ShouldBindJSON(&booking)
	log.Printf("%+v", booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	//tourが存在するか確認
	_, err = b.tourModelRepository.GetTour(booking.TourID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	//userが存在するか確認
	_, err = b.userModelRepository.Read(booking.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	UUID := uuid.New().String()
	booking.BookingID = UUID
	err = b.bookingModelRepository.Set(&booking)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (b BookingController) DeleteBooking(c *gin.Context) {
	bookingID := c.Param("bookingID")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_, err := b.bookingModelRepository.ReadByBookingID(bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	err = b.bookingModelRepository.Delete(bookingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (b BookingController) GetBookingByUserID(c *gin.Context) {
	bookingID := c.Param("userID")
	bookInfo, err := b.bookingModelRepository.ReadByUserID(bookingID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, bookInfo)
}
