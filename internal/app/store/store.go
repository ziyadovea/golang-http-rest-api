package store

// Store - интерфейс для хранилища
type Store interface {
	User() UserRepository
}
