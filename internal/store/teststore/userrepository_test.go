package teststore_test

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/store"
	"go-rest-api/internal/store/teststore"
	"testing"

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
	testEmail := "user@q.com"

	// check error for empty user
	_, err := s.User().FindByEmail(testEmail)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = testEmail

	// check correct selection
	s.User().Create(u)
	u, err = s.User().FindByEmail(testEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
