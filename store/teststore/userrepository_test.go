package teststore_test

import (
	"testing"

	"github.com/evdokimovsv/http-rest-api.git/model"
	"github.com/evdokimovsv/http-rest-api.git/store"
	"github.com/evdokimovsv/http-rest-api.git/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
