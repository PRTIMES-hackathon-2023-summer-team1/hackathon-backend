package models

import (
	"time"
)

type Booking struct {
	BookingID    string    `gorm:"primaryKey" json:"booking_id"`
	UserID       string    `gorm:"not null" json:"user_id"`
	PlanID       string    `gorm:"not null" json:"plan_id"`
	Participants int       `gorm:"not null" json:"participants"`
	Men          int       `json:"men"`
	Woman        int       `json:"woman"`
	Adults       int       `json:"adults"`
	Children     int       `json:"children"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
