package models

import (
	"time"
)

type User struct {
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique, not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
