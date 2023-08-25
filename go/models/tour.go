package models

import (
	"fmt"
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

// var citiesJSON = `
// [
//
//	  {
//	      "name": "東京",
//	      "description": "躍動する日本の首都",
//	      "body": "# 東京\n\n東京は日本の首都であり、国内最大の都市です。近代的な高層ビルと伝統的な文化が融合する、活気に満ちた都市です。\n\n## 観光名所\n\n東京タワーや浅草寺など、観光名所が豊富にあります。また、新宿や渋谷などの繁華街ではショッピングやエンターテインメントを楽しむことができます。\n\n## カルチャー\n\n歌舞伎や能などの伝統芸能から、最先端のファッションやアートまで、多様なカルチャーが交差する場所です。\n\n## 食文化\n\n寿司や刺身などの海産物から、ラーメンやお好み焼きまで、全国各地の料理が楽しめます。\n"
//	  },
//	  {
//	      "name": "京都",
//	      "description": "歴史と美の宝庫",
//	"body": "# 京都\n\n京都は日本の歴史と美が詰まった都市であり、多くの伝統的な寺社や庭園が点在しています。\n\n## 世界遺産\n\n清水寺や金閣寺など、多くの世界遺産が存在し、日本の歴史と仏教文化を感じることができます。\n\n## 着物文化\n\n観光客は着物を借りて街を歩くことができ、昔ながらの風情を楽しむことができます。\n\n## 茶道と料理\n\n抹茶を使った茶道や懐石料理など、伝統的な日本の文化と食を体験できる場所です。\:文化"
//	  },
//	  {
//	      "name": "大阪",
//	      "description": "食とショッピングの街",
//	      "body": "# 大阪\n\n大阪は食とショッピングが楽しめる都市として知られています。活気あふれる街並みが特徴です。\n\n## 道頓堀\n\n大阪のシンボル的なエリアであり、多彩な飲食店やショップが軒を連ね、活気溢れる雰囲気が漂っています。\n\n## グルメ\n\nたこ焼きやお好み焼きなど、大阪ならではの屋台料理が楽しめます。また、高級な食材を使ったレストランも充実しています。\n\n## ユニバーサル・スタジオ・ジャパン\n\n日本初のユニバーサル・スタジオがあり、映画の世界を楽しむことができます。\n"
//	  }
//
// ]
// `
type City struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

var cities = []City{
	{
		Name:        "東京",
		Description: "躍動する日本の首都",
		Body:        "# 東京\n\n東京は日本の首都であり、国内最大の都市です。近代的な高層ビルと伝統的な文化が融合する、活気に満ちた都市です。\n\n## 観光名所\n\n東京タワーや浅草寺など、観光名所が豊富にあります。また、新宿や渋谷などの繁華街ではショッピングやエンターテインメントを楽しむことができます。\n\n## カルチャー\n\n歌舞伎や能などの伝統芸能から、最先端のファッションやアートまで、多様なカルチャーが交差する場所です。\n\n## 食文化\n\n寿司や刺身などの海産物から、ラーメンやお好み焼きまで、全国各地の料理が楽しめます。\n",
	},
	{
		Name:        "京都",
		Description: "歴史と美の宝庫",
		Body:        "# 京都\n\n京都は日本の歴史と美が詰まった都市であり、多くの伝統的な寺社や庭園が点在しています。\n\n## 世界遺産\n\n清水寺や金閣寺など、多くの世界遺産が存在し、日本の歴史と仏教文化を感じることができます。\n\n## 着物文化\n\n観光客は着物を借りて街を歩くことができ、昔ながらの風情を楽しむことができます。\n\n## 茶道と料理\n\n抹茶を使った茶道や懐石料理など、伝統的な日本の文化と食を体験できる場所です。\n",
	},
	{
		Name:        "大阪",
		Description: "食とショッピングの街",
		Body:        "# 大阪\n\n大阪は食とショッピングが楽しめる都市として知られています。活気あふれる街並みが特徴です。\n\n## 道頓堀\n\n大阪のシンボル的なエリアであり、多彩な飲食店やショップが軒を連ね、活気溢れる雰囲気が漂っています。\n\n## グルメ\n\nたこ焼きやお好み焼きなど、大阪ならではの屋台料理が楽しめます。また、高級な食材を使ったレストランも充実しています。\n\n## ユニバーサル・スタジオ・ジャパン\n\n日本初のユニバーサル・スタジオがあり、映画の世界を楽しむことができます。\n",
	},
}

func getRandomCityInfo(cities []City) (string, string, string) {
	result := rand.Intn(2) + 1
	fmt.Println(result)
	fmt.Println(cities[result])
	randomCity := cities[result]
	return randomCity.Name, randomCity.Description, randomCity.Body
}

func NewDummyTour(userID string, isVisible bool, faker *faker.Faker) *Tour {

	jst, _ := time.LoadLocation("Asia/Tokyo")
	firstDay := faker.Time().Time(time.Now().AddDate(0, 0, 1)).In(jst)
	maxCapacity := faker.RandomDigitNotNull()

	name, description, body := getRandomCityInfo(cities)

	return &Tour{
		TourID:          faker.UUID().V4(),
		UserID:          userID,
		Name:            name,
		Description:     description,
		Body:            body,
		Price:           faker.RandomDigitNotNull(),
		FirstDay:        firstDay,
		LastDay:         firstDay.AddDate(0, 0, faker.RandomDigitNotNull()),
		MaxCapacity:     maxCapacity,
		CurrentCapacity: 0,
		IsVisible:       isVisible,
	}
}
