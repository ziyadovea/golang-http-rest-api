package sqlstore

import (
	"github.com/jackc/pgx/v4"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
)

// Store - структура для хранилища (БД)
type Store struct {
	connection     *pgx.Conn
	userRepository *UserRepository
}

// New создает экземпляр хранилища
func New(connection *pgx.Conn) *Store {
	return &Store{
		connection: connection,
	}
}

// User служит для создания репозитория
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}
