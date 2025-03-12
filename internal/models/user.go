package models

import (
	"github.com/google/uuid"
)

// @Description Структура, описывающая пользователя

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"` // Явно указываем UUID
	PhoneNumber  string    `json:"phone_number" validate:"required,e164"`
	EmailAdress  string    `json:"email_adress" validate:"required,email"`
	Password     string    `json:"password,omitempty" gorm:"-"` // `omitempty` убирает пустые значения из JSON
	PasswordHash string    `json:"-" gorm:"column:password_hash"`
	Name         string    `json:"name" validate:"required,min=3,max=50"`
}
