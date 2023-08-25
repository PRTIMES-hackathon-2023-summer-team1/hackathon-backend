package controllers

import (
	"errors"
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/gin-gonic/gin"
)

type TourController struct {
	tourRepository repository.ITourRepository
	userRepository repository.IUserRepository
}

func NewTourController(repo repository.ITourRepository, userRepo repository.IUserRepository) *TourController {
	return &TourController{tourRepository: repo, userRepository: userRepo}
}

func (t TourController) GetAllTours(c *gin.Context) {
	allTours, err := t.tourRepository.GetAllTours()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, allTours)
}

func (t TourController) GetTour(c *gin.Context) {
	tourId := c.Param("tour_id")
	if tourId == "" {
		err := errors.New("param is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	tourInfo, err := t.tourRepository.GetTour(tourId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, tourInfo)
}

func (t TourController) CreateTour(c *gin.Context) {
	var tour models.Tour
	userID, ok := c.Get("userID")
	if !ok {
		err := errors.New("userID is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	tour.UserID, ok = userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_, err := t.userRepository.ReadByID(tour.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	var tourInfoCreated models.Tour
	err = c.ShouldBindJSON(&tourInfoCreated)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	err = t.tourRepository.CreateTour(&tourInfoCreated)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, tourInfoCreated)
}

func (t TourController) EditTour(c *gin.Context) {
	var editedTourInfo models.Tour
	var tour models.Tour
	err := c.ShouldBindJSON(&editedTourInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		err := errors.New("userID is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	tour.UserID, ok = userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_, err = t.userRepository.ReadByID(tour.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	err = t.tourRepository.EditTour(editedTourInfo)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, editedTourInfo)
}

func (t TourController) SearchTour(c *gin.Context) {
	keyword := c.Query("keyword")
	result, err := t.tourRepository.SearchTour(keyword)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
