package model

import "testing"

// TestUser - вспомогательная функция, которая создает пользователя
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		ID:                0,
		Email:             "user@example.com",
		EncryptedPassword: "password",
	}
}
