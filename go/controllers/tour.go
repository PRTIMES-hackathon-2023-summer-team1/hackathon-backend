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
}

func NewTourController(repo repository.ITourRepository) *TourController {
	return &TourController{tourRepository: repo}
}

func (t TourController) GetAll(c *gin.Context) {
	allTours, err := t.tourRepository.GetAll()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, allTours)
}

func (t TourController) Get(c *gin.Context) {
	tourId := c.Param("tour_id")
	if tourId == "" {
		err := errors.New("Param is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	tourInfo, err := t.tourRepository.Get(tourId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, tourInfo)
	return
}

func (t TourController) CreateTour(c *gin.Context) {
	var tourInfoCreated models.Tour
	err := c.ShouldBindJSON(&tourInfoCreated)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	err = t.tourRepository.CreateTour(tourInfoCreated)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, tourInfoCreated)
}

func (t TourController) EditTour(c *gin.Context) {
	var editedTourInfo models.Tour
	err := c.ShouldBindJSON(&editedTourInfo)
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
