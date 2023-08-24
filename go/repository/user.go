package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(u models.User) error
	ReadByID(id string) (*models.User, error)
	ReadByEmail(email string) (*models.User, error)
	IsAdmin(userId string) (bool, error)
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

func (t UserRepository) ReadByID(id string) (*models.User, error) {
	user := &models.User{}
	result := t.repo.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (t UserRepository) ReadByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := t.repo.Where("email = ?", email).First(user)
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
	return user.IsAdmin, nil
}
