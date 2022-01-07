package store_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"testing"
)

// TestUserRepository_Create тестирует создание нового пользователя
func TestUserRepository_Create(t *testing.T) {
	st, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	err := st.User().Create(&model.User{
		ID:                0,
		Email:             "user1@example.com",
		EncryptedPassword: "",
	})
	assert.NoError(t, err)
}

// TestUserRepository_FindByEmail тестирует поиск пользователя по email
func TestUserRepository_FindByEmail(t *testing.T) {
	st, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	err := st.User().Create(&model.User{
		ID:                0,
		Email:             "user1@example.com",
		EncryptedPassword: "",
	})
	assert.NoError(t, err)

	// 1-ый кейс - ищем пользователя, которого нет
	u, err := st.User().FindByEmail("user2@example.com")
	assert.Error(t, err)
	assert.Nil(t, u)

	// 2-ой кейс - ищем существующего пользователя, все ок
	u, err = st.User().FindByEmail("user1@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
