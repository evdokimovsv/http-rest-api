package apiserver

import "testing"

import "net/http/httptest"

import "net/http"

import "github.com/evdokimovsv/http-rest-api.git/store/teststore"

import "github.com/stretchr/testify/assert"

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
