package models

import (
	"math/rand"
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

var japanPlaceNames = []string{
	"東京", "大阪", "京都", "札幌", "名古屋", "福岡", "広島", "仙台", "横浜", "神戸",
	"沖縄", "横浜", "川崎", "岡山", "新潟", "金沢", "長野", "静岡", "鹿児島", "宮崎",
	"熊本", "奈良", "長崎", "大分", "富山", "姫路", "岐阜", "滋賀", "宮城", "千葉",
	"埼玉", "群馬", "栃木", "茨城", "福島", "青森", "秋田", "山形", "岩手", "北海道",
}

func NewDummyTour(userID string, isVisible bool, faker *faker.Faker) *Tour {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	firstDay := faker.Time().Time(time.Now().AddDate(0, 0, 1)).In(jst)
	return &Tour{
		TourID:          faker.UUID().V4(),
		UserID:          userID,
		Name:            faker.Person().Name(),
		Description:     faker.Lorem().Sentence(10),
		Body:            selectRandomPlacename(),
		Price:           faker.RandomDigitNotNull(),
		FirstDay:        firstDay,
		LastDay:         firstDay.AddDate(0, 0, faker.RandomDigitNotNull()),
		MaxCapacity:     faker.RandomDigitNotNull(),
		CurrentCapacity: 0,
		IsVisible:       isVisible,
	}
}

func selectRandomPlacename() string {
	firstName := japanPlaceNames[rand.Intn(len(japanPlaceNames))]
	secondName := japanPlaceNames[rand.Intn(len(japanPlaceNames))]
	fullName := firstName + " " + secondName
	return fullName
}
