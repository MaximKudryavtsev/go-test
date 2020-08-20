package teststore_test

import (
	"fmt"
	"github.com/MaximKudryavtsev/go-test/internal/app/model"
	"github.com/MaximKudryavtsev/go-test/internal/app/store"
	"github.com/MaximKudryavtsev/go-test/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@user.com"
	_, err := s.User().FindByEmail(email)
	fmt.Println(err)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
