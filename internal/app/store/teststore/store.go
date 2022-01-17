package teststore

import (
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
)

// Store - структура для хранилища (БД)
type Store struct {
	userRepository *UserRepository
}

// New создает экземпляр хранилища
func New() *Store {
	return &Store{}
}

// User служит для создания репозитория
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int]*model.User),
		}
	}
	return s.userRepository
}
