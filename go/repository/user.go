package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/google/uuid"
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
	u.UserID = uuid.New().String() // UserIDはUUIDで生成
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
