package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/teststore"
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

	s := newServer(teststore.New())
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
