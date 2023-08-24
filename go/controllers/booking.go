package controllers

import (
	"errors"
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
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	//tourが存在するか確認
	_, err = b.tourModelRepository.GetTour(booking.TourID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, err.Error()})
		return
	}

	//userが存在するか確認
	_, err = b.userModelRepository.Read(booking.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, err.Error()})
		return
	}

	// participantsが1以上か確認
	if booking.Participants < 1 {
		err := errors.New("participants is less than 1")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	UUID := uuid.New().String()
	booking.BookingID = UUID
	err = b.bookingModelRepository.Set(&booking)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"booking_id": UUID, "user_id": booking.UserID, "tour_id": booking.TourID, "participants": booking.Participants})
}

func (b BookingController) DeleteBooking(c *gin.Context) {
	bookingID := c.Param("bookingID")
	// bookingIDが空か確認
	if bookingID == "" {
		err := errors.New("bookingID is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// bookingが存在するか確認
	_, err := b.bookingModelRepository.ReadByBookingID(bookingID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, err.Error()})
		return
	}

	// bookingを削除
	err = b.bookingModelRepository.Delete(bookingID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "booking deleted"})
}

func (b BookingController) GetBookingByUserID(c *gin.Context) {
	userID := c.Param("userID")
	bookInfo, err := b.bookingModelRepository.ReadByUserID(userID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookInfo)
}
