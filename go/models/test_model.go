package models

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/db"
)

type Test struct {
	UID  string `json:"uid" gorm:"primaryKey"`
	Name string `json:"name"`
}

type TestModel struct{}

func (t TestModel) Set(te Test) error {
	err := db.DB.Save(&te).Where("uid = ?", te.UID).Error
	if err != nil {
		return err
	}
	return nil
}
