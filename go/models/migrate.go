package models

import (
	"github.com/jaswdr/faker"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

const (
	USERS_COUNT       = 30
	ADMIN_USERS_COUNT = 5
	VISIALBE_TOURS    = 30
	UNVISIALBE_TOURS  = 5
	BOOKING_COUNT     = 50
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&TestModel{}, &User{}, &Booking{}, &Tour{})
}

func Drop(db *gorm.DB) {
	db.Migrator().DropTable(&TestModel{}, &User{}, &Booking{}, &Tour{})
}

func InsertDummyDatas(db *gorm.DB) {
	faker := faker.New()
	users := []User{}
	for i := 0; i < USERS_COUNT; i++ {
		users = append(users, *NewDummyUser(false, &faker))
	}
	db.Create(&users)

	adminUsers := []User{}
	for i := 0; i < ADMIN_USERS_COUNT; i++ {
		adminUsers = append(adminUsers, *NewDummyUser(true, &faker))
	}
	db.Create(&adminUsers)

	visiableTours := []Tour{}
	for i := 0; i < VISIALBE_TOURS; i++ {
		userID := rand.Intn(ADMIN_USERS_COUNT)
		visiableTours = append(visiableTours, *NewDummyTour(users[userID].UserID, true, &faker))
	}
	db.Create(&visiableTours)

	unvisiableTours := []Tour{}
	for i := 0; i < UNVISIALBE_TOURS; i++ {
		userID := rand.Intn(ADMIN_USERS_COUNT)
		unvisiableTours = append(unvisiableTours, *NewDummyTour(users[userID].UserID, false, &faker))
	}
	db.Create(&unvisiableTours)

	bookings := []Booking{}
	for i := 0; i < BOOKING_COUNT; i++ {
		userID := rand.Intn(USERS_COUNT)
		tourID := rand.Intn(VISIALBE_TOURS)
		bookings = append(bookings, *NewDummyBooking(users[userID].UserID, visiableTours[tourID].TourID, &faker))
	}
	db.Create(&bookings)
}
