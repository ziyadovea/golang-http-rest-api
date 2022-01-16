package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// User - структура пользователя в БД
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email" validate:"required,email"`
	EncryptedPassword string `json:"-" validate:"required,gte=6,lte=100"`
}

var validate *validator.Validate

// Validate проверяет валидность полей пользователя
func (u *User) Validate() error {
	validate = validator.New()
	return validate.Struct(u)
}

// ComparePassword сравнивает пароли
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}
