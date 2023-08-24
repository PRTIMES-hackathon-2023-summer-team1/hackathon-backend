package repository

import (
	"log"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IBookingRepository interface {
	Set(booking *models.Booking) error
	Delete(bookingID string) error
	ReadByUserID(userID string) ([]*models.Booking, error)
	ReadByBookingID(bookingID string) (*models.Booking, error)
}

type BookingRepository struct {
	repo *gorm.DB
}

func NewBookingRepository(repo *gorm.DB) *BookingRepository {
	return &BookingRepository{repo: repo}
}

func (b BookingRepository) Set(booking *models.Booking) error {
	var tour *models.Tour
	err := b.repo.Model(&tour).Where("tour_id = ?", booking.TourID).First(&tour).Error
	if err != nil {
		return err
	}

	if tour.CurrentCapacity + booking.Participants > tour.MaxCapacity {
		return err
	}

	tx := b.repo.Begin()
	{
		err := tx.Create(&booking).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		tour.CurrentCapacity += booking.Participants
		err = b.repo.Save(&tour).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit().Error
	return err
}

func (b BookingRepository) Delete(bookingID string) error {
	var booking *models.Booking
	err := b.repo.Where("booking_id = ?", bookingID).First(&booking).Error
	if err != nil {
		return err
	}
	tx := b.repo.Begin()
	{
		err = tx.Delete(&booking).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		log.Println(booking)
		var tour *models.Tour
		err = b.repo.Model(&tour).Where("tour_id = ?", booking.TourID).First(&tour).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		tour.CurrentCapacity += booking.Participants
		err = b.repo.Save(&tour).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit().Error
	return err
}

func (b BookingRepository) ReadByUserID(userID string) ([]*models.Booking, error) {
	var booking []*models.Booking
	// err := b.repo.Where("user_id = ?", userID).Preload("Tours").Find(&booking).Error
	err := b.repo.Preload("Tour").Where("user_id = ?", userID).Find(&booking).Error
	return booking, err
}

func (b BookingRepository) ReadByBookingID(bookingID string) (*models.Booking, error) {
	var booking *models.Booking
	err := b.repo.Where("booking_id = ?", bookingID).First(&booking).Error
	return booking, err
}
