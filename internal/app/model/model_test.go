package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"strings"
	"testing"
)

// TestUser_Validate тестирует валидацию полей пользователя
func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name       string
		createUser func() *model.User
		isValid    bool
	}{
		{
			"Валидный кейс",
			func() *model.User {
				u := model.TestUser(t)
				return u
			},
			true,
		},
		{
			"Пустой email",
			func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			false,
		},
		{
			"Некорректный email",
			func() *model.User {
				u := model.TestUser(t)
				u.Email = "incorrectEmail"
				return u
			},
			false,
		},
		{
			"Пустой пароль",
			func() *model.User {
				u := model.TestUser(t)
				u.EncryptedPassword = ""
				return u
			},
			false,
		},
		{
			"Слишком маленький пароль",
			func() *model.User {
				u := model.TestUser(t)
				u.EncryptedPassword = "pass"
				return u
			},
			false,
		},
		{
			"Слишком длинный пароль",
			func() *model.User {
				u := model.TestUser(t)
				u.EncryptedPassword = strings.Repeat("p", 101)
				return u
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.createUser().Validate())
			} else {
				assert.Error(t, tc.createUser().Validate())
			}
		})
	}
}
