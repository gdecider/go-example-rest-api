package store_test

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	// databaseURL - global variable, setted in "store_test.go" or from env DATABASE_URL
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	// databaseURL - global variable, setted in "store_test.go" or from env DATABASE_URL
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	testEmail := "user@q.com"

	// check error for empty user
	_, err := s.User().FindByEmail(testEmail)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = testEmail

	// check correct selection
	s.User().Create(u)
	u, err = s.User().FindByEmail(testEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
