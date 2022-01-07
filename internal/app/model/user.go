package model

// User - структура пользователя в БД
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
