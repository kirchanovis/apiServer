package teststore_test

import (
	"testing"

	"github.com/kirchanovis/apiServer/internal/app/model"
	"github.com/kirchanovis/apiServer/internal/app/store"
	"github.com/kirchanovis/apiServer/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRrepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRrepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	email := u.Email
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
