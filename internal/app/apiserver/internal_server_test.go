package apiserver

import (
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleHello(t *testing.T) {
	server := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	assert.NoError(t, err)
	server.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "Hello!", rec.Body.String())
}
