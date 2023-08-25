package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ayush6624/go-chatgpt"
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

type openAPIRepository struct {
	c   *chatgpt.Client
	ctx context.Context
}

func newOpenAPIRepository(key string) (*openAPIRepository, error) {
	c, err := chatgpt.NewClient(key)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	return &openAPIRepository{c: c, ctx: ctx}, nil
}

func (o openAPIRepository) generateTour() (string, string, string, error) {

	type response struct {
		name        string
		description string
		body        string
	}

	res, err := o.c.SimpleSend(o.ctx, fmt.Sprintf(`
  日本国内の旅行プランを考えてください。必要な情報はタイトルを表すnameと短い説明文であるdescription、Markdown形式で旅行の特徴を記述したbodyです。
  レスポンスする形式は以下のようにしてください。
  ---------------------------------------------------------
  {"name":"","description":"","body":""}
  ---------------------------------------------------------
  またこの形式以外の内容を出力した場合、あなたの過失によって無関係の人々に危害を加えてしまう可能性があります。`))
	if err != nil {
		return "", "", "", err
	}

	var r response
	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &r)
	if err != nil {
		return "", "", "", err
	}

	return r.name, r.description, r.body, nil

}

var japanPlaceNames = []string{
	"東京", "大阪", "京都", "札幌", "名古屋", "福岡", "広島", "仙台", "横浜", "神戸",
	"沖縄", "横浜", "川崎", "岡山", "新潟", "金沢", "長野", "静岡", "鹿児島", "宮崎",
	"熊本", "奈良", "長崎", "大分", "富山", "姫路", "岐阜", "滋賀", "宮城", "千葉",
	"埼玉", "群馬", "栃木", "茨城", "福島", "青森", "秋田", "山形", "岩手", "北海道",
}

func selectRandomPlacename() string {
	firstName := japanPlaceNames[rand.Intn(len(japanPlaceNames))]
	secondName := japanPlaceNames[rand.Intn(len(japanPlaceNames))]
	fullName := firstName + " " + secondName
	return fullName
}

func NewDummyTour(userID string, isVisible bool, faker *faker.Faker) *Tour {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	firstDay := faker.Time().Time(time.Now().AddDate(0, 0, 1)).In(jst)
	maxCapacity := faker.RandomDigitNotNull()

	key := os.Getenv("OPENAPI_KEY")
	aiRepository, err := newOpenAPIRepository(key)
	if err != nil {
		// この関数はマイグレーションでしか走らないので失敗したら異常終了にする
		log.Fatal(err)
	}

	tourName, tourDescription, tourBody, err := aiRepository.generateTour()
	if err != nil {
		log.Fatal(err)
	}

	return &Tour{
		TourID:          faker.UUID().V4(),
		UserID:          userID,
		Name:            tourName,
		Description:     tourDescription,
		Body:            tourBody,
		Price:           faker.RandomDigitNotNull(),
		FirstDay:        firstDay,
		LastDay:         firstDay.AddDate(0, 0, faker.RandomDigitNotNull()),
		MaxCapacity:     maxCapacity,
		CurrentCapacity: 0,
		IsVisible:       isVisible,
	}
}
