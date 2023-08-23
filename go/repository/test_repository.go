package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type ITestRepository interface {
	Set(t models.TestModel) error
}

type TestRepository struct {
	repo *gorm.DB
}

func NewTestRepository(repo *gorm.DB) *TestRepository {
	return &TestRepository{repo: repo}
}

func (t TestRepository) Set(testInfo models.TestModel) error {
	err := t.repo.Create(&testInfo).Error
	return err
}
