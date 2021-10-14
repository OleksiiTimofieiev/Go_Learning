package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user3@example.org",
		Password: "password",
	}
}
