package repository

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"gorm.io/gorm"
)

type IJTIRepository interface {
	Create(JWTJTI *models.JWTJTI) error
	Delete(jti string) error
	IsValid(jti string) (bool, error)
}

type JTIRepository struct {
	repo *gorm.DB
}

func NewJTIRepository(repo *gorm.DB) *JTIRepository {
	return &JTIRepository{repo: repo}
}

func (t *JTIRepository) Create(JWTJTI *models.JWTJTI) error {
	result := t.repo.Create(&JWTJTI)
	return result.Error
}

func (t *JTIRepository) Delete(jti string) error {
	result := t.repo.Where("jti = ?", jti).Delete(&models.JWTJTI{})
	return result.Error
}

func (t *JTIRepository) IsValid(jti string) (bool, error) {
	var JWTJTI *models.JWTJTI
	err := t.repo.Where("jti = ?", jti).First(&JWTJTI).Error
	if err != nil {
		return false, err
	}
	if JWTJTI == nil {
		return false, nil
	}
	return JWTJTI.IsValid(), nil
}
