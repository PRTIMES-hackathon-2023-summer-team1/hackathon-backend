package models

import (
	"time"

	"github.com/jaswdr/faker"
)

type Booking struct {
	BookingID    string    `gorm:"primaryKey" json:"booking_id"`
	UserID       string    `gorm:"not null" json:"user_id"`
	TourID       string    `gorm:"not null" json:"tour_id"`
	Participants int       `gorm:"not null" json:"participants"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewDummyBooking(userID string, tourID string, faker *faker.Faker) *Booking {
	return &Booking{
		BookingID:    faker.UUID().V4(),
		UserID:       userID,
		TourID:       tourID,
		Participants: faker.RandomDigitNotNull(),
	}
}
