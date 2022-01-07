package store

import (
	"context"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
)

// UserRepository - структура для репозитория для БД
type UserRepository struct {
	store *Store
}

// Create создает пользователя
func (r *UserRepository) Create(u *model.User) error {
	err := r.store.connection.QueryRow(
		context.Background(),
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	return err
}

// FindByEmail ищет пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
