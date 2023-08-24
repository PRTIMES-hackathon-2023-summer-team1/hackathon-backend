package repository

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITourRepository interface {
	GetAllTours() ([]models.Tour, error)
	GetTour(string) (models.Tour, error)
	CreateTour(*models.Tour) error
	EditTour(models.Tour) error
	SearchTour(keyword string) ([]models.Tour, error)
}

type TourRepository struct {
	repo *gorm.DB
}

func NewTourRepository(repo *gorm.DB) *TourRepository {
	return &TourRepository{repo: repo}
}

func (t TourRepository) GetAllTours() ([]models.Tour, error) {
	var allTours []models.Tour
	err := t.repo.Find(&allTours).Error
	if err != nil {
		return nil, err
	}
	return allTours, nil
}

func (t TourRepository) GetTour(tourId string) (models.Tour, error) {
	var tourInfo models.Tour
	err := t.repo.First(&tourInfo, "tour_id = ?", tourId).Error
	if err != nil {
		return models.Tour{}, err
	}
	return tourInfo, nil
}

func (t TourRepository) CreateTour(to *models.Tour) error {
	//送られてきた時間がRFC3339か確認
	firstDay := to.FirstDay.Format(time.RFC3339)
	_, err := time.Parse(time.RFC3339, firstDay)
	if err != nil {
		return err
	}

	lastDay := to.FirstDay.Format(time.RFC3339)
	_, err = time.Parse(time.RFC3339, lastDay)
	if err != nil {
		return err
	}

	_, err = time.Parse(time.RFC3339, to.LastDay.String())
	to.TourID = uuid.New().String()
	err = t.repo.Omit("CreatedAt", "UpdatedAt").Create(&to).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TourRepository) EditTour(to models.Tour) error {
	//送られてきた時間がRFC3339か確認
	firstDay := to.FirstDay.Format(time.RFC3339)
	_, err := time.Parse(time.RFC3339, firstDay)
	if err != nil {
		return err
	}

	lastDay := to.FirstDay.Format(time.RFC3339)
	_, err = time.Parse(time.RFC3339, lastDay)
	if err != nil {
		return err
	}
	err = t.repo.Model(&models.Tour{}).Where("tour_id = ?", to.TourID).Updates(map[string]interface{}{
		"Description":     to.Description,
		"Body":            to.Body,
		"Price":           to.Price,
		"FirstDay":        to.FirstDay,
		"LastDay":         to.LastDay,
		"MaxCapacity":     to.MaxCapacity,
		"CurrentCapacity": to.CurrentCapacity,
		"IsVisible":       to.IsVisible,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TourRepository) SearchTour(keyword string) ([]models.Tour, error) {
	parsedKeyword, err := url.QueryUnescape(keyword)
	if err != nil {
		return nil, err
	}
	middleKeywords := strings.Replace(parsedKeyword, "　", " ", -1)
	keywords := strings.Split(middleKeywords, " ")
	condition := "body LIKE ?"
	values := []interface{}{"%" + keywords[0] + "%"}

	for i := 1; i < len(keywords); i++ {
		condition += " AND body LIKE ?"
		values = append(values, "%"+keywords[i]+"%")
	}

	var tours []models.Tour
	err = t.repo.Where(condition, values...).Find(&tours).Error
	if err != nil {
		return nil, err
	}
	return tours, nil
}
