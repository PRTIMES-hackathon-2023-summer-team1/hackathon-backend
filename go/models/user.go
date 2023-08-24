package models

import (
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/utility"
	"github.com/jaswdr/faker"
)

type User struct {
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDummyUser(isAdmin bool, faker *faker.Faker) *User {
	encryptedPassword, _ := utility.EncryptPassword(faker.Internet().Password())
	return &User{
		UserID:   faker.UUID().V4(),
		Name:     faker.Person().Name(),
		Email:    faker.Internet().Email(),
		Password: encryptedPassword,
		IsAdmin:  isAdmin,
	}
}
