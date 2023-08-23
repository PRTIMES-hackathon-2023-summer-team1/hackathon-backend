package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITourRepository interface {
	GetAll() (models.Tour, error)
	Get(string) (models.Tour, error)
	CreateTour(models.Tour) error
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

func (t TourRepository) CreateTour(to models.Tour) error {
	//newUUID := uuid.New()
	//送られてきた時間形式がRFC3339かバリデーションを行う
	//	_, err := time.Parse(time.RFC3339, strings.Join(strings.Fields(to.FirstDay.String()), ""))
	//	if err != nil {
	//		return err
	//	}
	//	_, err = time.Parse(time.RFC3339, strings.Join(strings.Fields(to.LastDay.String()), ""))
	//	if err != nil {
	//		return err
	//	}
	to.TourID = uuid.New().String()
	err := t.repo.Omit("CreatedAt", "UpdatedAt").Create(&to).Error
	if err != nil {
		return err
	}
	return nil
}
