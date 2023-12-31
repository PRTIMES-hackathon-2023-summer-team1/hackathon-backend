package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IBookingRepository interface {
	Set(booking *models.Booking) error
	Delete(bookingID string) error
	ReadByUserID(userID string) ([]*models.BookingJoinTour, error)
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
	tx := b.repo.Begin()
	{
		err := b.repo.Model(&tour).Where("tour_id = ?", booking.TourID).First(&tour).Error
		if err != nil {
			return err
		}

		if tour.CurrentCapacity+booking.Participants > tour.MaxCapacity {
			return err
		}

		err = tx.Create(&booking).Error
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
	err := tx.Commit().Error
	return err
}

func (b BookingRepository) Delete(bookingID string) error {
	var booking *models.Booking
	tx := b.repo.Begin()
	{
		err := b.repo.Where("booking_id = ?", bookingID).First(&booking).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&booking).Error
		if err != nil {
			tx.Rollback()
			return err
		}

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
	err := tx.Commit().Error
	return err
}

func (b BookingRepository) ReadByUserID(userID string) ([]*models.BookingJoinTour, error) {
	var bookingJoinTour []*models.BookingJoinTour
	err := b.repo.Model(&models.Booking{}).Select("bookings.booking_id, bookings.tour_id, bookings.user_id, tours.name, bookings.participants, tours.price, tours.first_day, tours.last_day").Joins("left join tours on bookings.tour_id = tours.tour_id").Where("bookings.user_id = ?", userID).Find(&bookingJoinTour).Error
	return bookingJoinTour, err
}

func (b BookingRepository) ReadByBookingID(bookingID string) (*models.Booking, error) {
	var booking *models.Booking
	err := b.repo.Where("booking_id = ?", bookingID).First(&booking).Error
	return booking, err
}
