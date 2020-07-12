package store_test

import (
	"testing"

	"github.com/kirchanovis/apiServer/internal/app/model"
	"github.com/kirchanovis/apiServer/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRrepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@mail.ru",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRrepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@mail.ru",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
