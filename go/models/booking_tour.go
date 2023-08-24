package models

import "time"

type BookingJoinTour struct {
	BookingID 		 		string    `json:"booking_id"`
	TourID            string    `json:"tour_id"`
	UserID 						string    `json:"user_id"`
	Name 					    string    `json:"name"`
	Participants      int       `json:"participants"`
	Price 					  int       `json:"price"`
	FirstDay          time.Time `json:"first_day"`
	LastDay           time.Time `json:"last_day"`
}
