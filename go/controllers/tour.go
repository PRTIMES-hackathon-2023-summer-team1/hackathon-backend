package controllers

import (
	"net/http"

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
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, allTours)
}
