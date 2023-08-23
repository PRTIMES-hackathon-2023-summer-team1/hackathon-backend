package models

import "time"

type Tour struct {
	TourID          string    `gorm:"primaryKey" json:"tour_id"`
	UserID          string    `gorm:"not null" json:"user_id"`
	Name            string    `gorm:"not null" json:"name"`
	Description     string    `json:"description"`
	Body            string    `json:"body"`
	Price           int       `json:"price"`
	FirstDay        time.Time `json:"first_day"`
	LastDay         time.Time `json:"last_day"`
	MaxCapacity     int       `json:"max_capacity"`
	CurrentCapacity int       `json:"current_capacity"`
	IsVisible       bool      `json:"is_visible"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}