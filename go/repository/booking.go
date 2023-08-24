package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IBookingRepository interface {
	Set(booking *models.Booking) error
	Delete(bookingID string) error
	ReadByUserID(userID string) ([]models.Booking, error)
	ReadByBookingID(bookingID string) (models.Booking, error)
}

type BookingRepository struct {
	repo *gorm.DB
}

func NewBookingRepository(repo *gorm.DB) *BookingRepository {
	return &BookingRepository{repo: repo}
}

func (b BookingRepository) Set(booking *models.Booking) error {
	err := b.repo.Create(&booking).Error
	return err
}

func (b BookingRepository) Delete(bookingID string) error {
	err := b.repo.Delete(&models.Booking{}, bookingID).Error
	return err
}

func (b BookingRepository) ReadByUserID(userID string) ([]models.Booking, error) {
	var booking []models.Booking
	// err := b.repo.Where("user_id = ?", userID).Preload("Tours").Find(&booking).Error
	err := b.repo.Preload("Tour").Where("user_id = ?", userID).Find(&booking).Error
	return booking, err
}

func (b BookingRepository) ReadByBookingID(bookingID string) (models.Booking, error) {
	var booking models.Booking
	err := b.repo.Where("tour_id = ?", bookingID).First(&booking).Error
	return booking, err
}
