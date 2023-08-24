package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(u models.User) error
	Read(id string) (*models.User, error)
	IsAdmin(string) (bool, error)
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
	result := t.repo.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (t UserRepository) IsAdmin(userId string) (bool, error) {
	var user models.User
	err := t.repo.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return false, err
	}
	if !user.IsAdmin {
		return false, nil
	}
	return true, nil
}
