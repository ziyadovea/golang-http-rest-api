package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/teststore"
	"testing"
)

// TestUserRepository_Create тестирует создание нового пользователя
func TestUserRepository_Create(t *testing.T) {
	store := teststore.New()
	repo := store.User()
	error := repo.Create(model.TestUser(t))
	assert.NoError(t, error)
}

// TestUserRepository_FindByID тестирует поиск пользователя по email
func TestUserRepository_FindByID(t *testing.T) {

	st := teststore.New()
	repo := st.User()
	u := model.TestUser(t)
	error := repo.Create(u)
	assert.NoError(t, error)

	// 1-ый кейс - ищем пользователя, которого нет
	user, err := repo.FindByID(1000)
	assert.Error(t, err)
	assert.EqualError(t, err, store.ErrUserNotFound.Error())
	assert.Nil(t, user)

	// 2-ой кейс - ищем существующего пользователя, все ок
	user, err = repo.FindByID(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)

}

// TestUserRepository_FindByEmail тестирует поиск пользователя по email
func TestUserRepository_FindByEmail(t *testing.T) {

	st := teststore.New()
	repo := st.User()
	u := model.TestUser(t)
	u.Email = "user1@example.com"
	error := repo.Create(u)
	assert.NoError(t, error)

	// 1-ый кейс - ищем пользователя, которого нет
	user, err := repo.FindByEmail("user2@example.com")
	assert.Error(t, err)
	assert.EqualError(t, err, store.ErrUserNotFound.Error())
	assert.Nil(t, user)

	// 2-ой кейс - ищем существующего пользователя, все ок
	user, err = repo.FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

}
