package main

import (
	"context"
	"testing"

	"github.com/Fenroe/carbonarapi/internal/database"
)

func TestCreateUser(t *testing.T) {
	queries, cleanup := dbTestSetup()
	defer cleanup()
	params := database.CreateUserParams{
		Email:          "testing@carbonara.com",
		HashedPassword: "myhashedpassword",
	}
	result, err := queries.CreateUser(context.Background(), params)
	if err != nil {
		t.Fatalf("could not create user: %s", err)
	}
	if result.Email != "testing@carbonara.com" {
		t.Fatalf("mismatch on Email field; expected 'testing@carbonara.com', but got '%s'", result.Email)
	}
}
