package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/teststore"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlerUsersCreate(t *testing.T) {

	testCases := []struct {
		name    string
		content interface{}
		code    int
	}{
		{
			name: "valid",
			content: map[string]string{
				"email":    "user@example.com",
				"password": "password",
			},
			code: http.StatusCreated,
		},
		{
			name:    "invalid request data",
			content: "invalid",
			code:    http.StatusBadRequest,
		},
		{
			name: "invalid fields",
			content: map[string]string{
				"email":    "user",
				"password": "",
			},
			code: http.StatusUnprocessableEntity,
		},
	}

	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("key")))
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			body := &bytes.Buffer{}
			err := json.NewEncoder(body).Encode(tc.content)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/users", body)
			s.handlerUsersCreate().ServeHTTP(rec, req)
		})
	}
}

func Test_handlerSessionsCreate(t *testing.T) {

	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("key")))
	password := "password"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	u := model.TestUser(t)
	u.EncryptedPassword = string(hashPassword)
	err = s.store.User().Create(u)
	assert.NoError(t, err)

	testCases := []struct {
		name           string
		content        interface{}
		expectedStatus int
	}{
		{
			name: "valid",
			content: map[string]string{
				"email":    "user@example.com",
				"password": "password",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid request",
			content:        "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "incorrect email",
			content: map[string]string{
				"email":    "incorrect@example.com",
				"password": "password",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "incorrect password",
			content: map[string]string{
				"email":    "user@example.com",
				"password": "incorrect",
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			body := &bytes.Buffer{}
			err = json.NewEncoder(body).Encode(tc.content)
			assert.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/sessions", body)
			s.handlerSessionsCreate().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
