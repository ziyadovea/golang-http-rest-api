package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/sqlstore"
	"testing"
)

// TestUserRepository_Create тестирует создание нового пользователя
func TestUserRepository_Create(t *testing.T) {
	conn, tearDown := sqlstore.TestConnection(t, databaseURL)
	defer tearDown("users")

	st := sqlstore.New(conn)
	err := st.User().Create(model.TestUser(t))
	assert.NoError(t, err)
}

// TestUserRepository_FindByEmail тестирует поиск пользователя по email
func TestUserRepository_FindByEmail(t *testing.T) {
	conn, tearDown := sqlstore.TestConnection(t, databaseURL)
	defer tearDown("users")

	st := sqlstore.New(conn)
	user := model.TestUser(t)
	user.Email = "user1@example.com"
	err := st.User().Create(user)
	assert.NoError(t, err)

	// 1-ый кейс - ищем пользователя, которого нет
	u, err := st.User().FindByEmail("user2@example.com")
	assert.Error(t, err)
	assert.EqualError(t, err, store.ErrUserNotFound.Error())
	assert.Nil(t, u)

	// 2-ой кейс - ищем существующего пользователя, все ок
	u, err = st.User().FindByEmail("user1@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
