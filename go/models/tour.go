package models

import (
	"time"

	"github.com/jaswdr/faker"
)

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

func NewDummyTour(userID string, isVisible bool, faker *faker.Faker) *Tour {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	firstDay := faker.Time().Time(time.Now().AddDate(0, 0, 1)).In(jst)
	maxCapacity := faker.RandomDigitNotNull()
	return &Tour{
		TourID:          faker.UUID().V4(),
		UserID:          userID,
		Name:            faker.Person().Name(),
		Description:     faker.Lorem().Sentence(10),
		Body:            faker.Lorem().Paragraph(5),
		Price:           faker.RandomDigitNotNull(),
		FirstDay:        firstDay,
		LastDay:         firstDay.AddDate(0, 0, faker.RandomDigitNotNull()),
		MaxCapacity:     maxCapacity,
		CurrentCapacity: 0,
		IsVisible:       isVisible,
	}
}
