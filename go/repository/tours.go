package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type ITourRepository interface {
	GetAll() (models.Tour, error)
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
		return allTours, err
	}
	return allTours, nil
}
