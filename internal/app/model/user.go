package model

import "github.com/go-playground/validator/v10"

// User - структура пользователя в БД
type User struct {
	ID                int
	Email             string `validate:"required,email"`
	EncryptedPassword string `validate:"required,gte=6,lte=100"`
}

var validate *validator.Validate

// Validate проверяет валидность полей пользователя
func (u *User) Validate() error {
	validate = validator.New()
	return validate.Struct(u)
}
