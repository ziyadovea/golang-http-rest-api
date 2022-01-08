package teststore

import (
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
)

// UserRepository - структура для репозитория для БД
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create создает пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}

// FindByEmail ищет пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	if u, ok := r.users[email]; ok {
		return u, nil
	}
	return nil, store.ErrUserNotFound
}
