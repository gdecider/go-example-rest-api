package sqlstore_test

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/store"
	"go-rest-api/internal/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	// databaseURL - global variable, setted in "store_test.go" or from env DATABASE_URL
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	// databaseURL - global variable, setted in "store_test.go" or from env DATABASE_URL
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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

func TestUserRepository_Find(t *testing.T) {
	// databaseURL - global variable, setted in "store_test.go" or from env DATABASE_URL
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
