package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@mail.ru",
		Password: "password",
	}
}
