package models

import (
	"time"
)

type Booking struct {
	BookingID    string    `gorm:"primaryKey" json:"booking_id"`
	UserID       string    `gorm:"not null" json:"user_id"`
	TourID       string    `gorm:"not null" json:"tour_id"`
	Participants int       `gorm:"not null" json:"participants"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
