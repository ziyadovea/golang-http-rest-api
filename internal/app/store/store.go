package store

import (
	"context"
	"github.com/jackc/pgx/v4"
)

// Store - структура для хранилища (БД)
type Store struct {
	config         *Config
	connection     *pgx.Conn
	userRepository *UserRepository
}

// New создает экземпляр хранилища
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open открывает коннект к БД
func (s *Store) Open() error {

	conn, err := pgx.Connect(context.Background(), s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// Пинг для проверки соединения
	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	s.connection = conn
	return nil

}

// Close закрывает коннект к БД
func (s *Store) Close() {
	s.connection.Close(context.Background())
}

// User служит для создания репозитория
func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}
