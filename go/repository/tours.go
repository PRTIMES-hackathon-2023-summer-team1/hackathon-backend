package repository

import (
	"fmt"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type ITourRepository interface {
	GetAll() (models.Tour, error)
	Get(string) (models.Tour, error)
}

type TourRepository struct {
	repo *gorm.DB
}

func NewTourRepository(repo *gorm.DB) *TourRepository {
	return &TourRepository{repo: repo}
}

func (t TourRepository) GetAll() (models.Tour, error) {
	var allTours models.Tour
	err := t.repo.Find(&allTours).Error
	if err != nil {
		return models.Tour{}, err
	}
	return allTours, nil
}

func (t TourRepository) Get(tourId string) (models.Tour, error) {
	var tourInfo models.Tour
	err := t.repo.First(&tourInfo, "tour_id = ?", tourId).Error
	if err != nil {
		return models.Tour{}, err
	}
	return tourInfo, nil
}
