package teststore

import (
	"errors"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
)

// UserRepository - структура для репозитория для БД
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create создает пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	u.ID = len(r.users) + 1
	if _, isExist := r.users[u.ID]; isExist {
		return errors.New("пользователь с таким ID уже существует")
	}
	r.users[u.ID] = u
	return nil
}

// FindByID ищет пользователя по email
func (r *UserRepository) FindByID(ID int) (*model.User, error) {
	if u, isExist := r.users[ID]; isExist {
		return u, nil
	}
	return nil, store.ErrUserNotFound
}

// FindByEmail ищет пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, v := range r.users {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, store.ErrUserNotFound
}
