package models

type TestModel struct {
	UID  string `json:"uid" gorm:"primaryKey"`
	Name string `json:"name"`
}
