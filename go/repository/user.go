package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(u models.User) error
	Read(id string) (*models.User, error)
}

type UserRepository struct {
	repo *gorm.DB
}

func NewUserRepository(repo *gorm.DB) *UserRepository {
	return &UserRepository{repo: repo}
}

func (t UserRepository) Create(u models.User) error {
	result := t.repo.Create(&u)
	return result.Error
}

func (t UserRepository) Read(id string) (*models.User, error) {
	user := &models.User{}
	result := t.repo.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
