package repository

import (
	"gorm.io/gorm"

	"backend/internal/models"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(client *models.User) error {
	return r.db.Table("users").Create(client).Error
}

func (r *AuthRepository) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var client models.User
	err := r.db.Table("users").
		Where("phone_number = ? OR phone_number = ?", phoneNumber, phoneNumber[1:]).
		First(&client).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &client, nil
}
