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

	u, err := s.User().Create(&model.User{
		Email: "user@q.com",
	})

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

	// check correct selection
	s.User().Create(&model.User{
		Email: testEmail,
	})
	u, err := s.User().FindByEmail(testEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
