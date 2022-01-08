package store

import "github.com/ziyadovea/golang-http-rest-api/internal/app/model"

// UserRepository - интерфейс для репозитория пользователя
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
