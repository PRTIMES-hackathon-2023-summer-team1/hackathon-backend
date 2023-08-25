package models

import (
	"time"

	"github.com/jaswdr/faker"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(faker.Internet().Password()), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &User{
		UserID:   faker.UUID().V4(),
		Name:     faker.Person().Name(),
		Email:    faker.Internet().Email(),
		Password: string(hash),
		IsAdmin:  isAdmin,
	}
}
