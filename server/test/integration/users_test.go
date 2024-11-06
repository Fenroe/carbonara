package test

import (
	"context"
	"testing"

	"github.com/Fenroe/carbonarapi/internal/database"
)

func TestUsers(t *testing.T) {
	// setup and cleanup
	queries, cleanup := dbTestSetup()
	defer cleanup()
	// create user
	createUserParams := database.CreateUserParams{
		Email:          "testing@carbonara.com",
		HashedPassword: "myhashedpassword",
	}
	createUserRes, err := queries.CreateUser(context.Background(), createUserParams)
	if err != nil {
		t.Fatalf("could not create user: %s", err)
	}
	if createUserRes.Email != "testing@carbonara.com" {
		t.Fatalf("mismatch on Email field; expected 'testing@carbonara.com', but got '%s'", createUserRes.Email)
	}
	// email field is unique
	_, err = queries.CreateUser(context.Background(), createUserParams)
	if err == nil {
		t.Fatal("duplicate user was created")
	}
	// get user by email
	getUserRes, err := queries.GetUserByEmail(context.Background(), createUserRes.Email)
	if err != nil {
		t.Fatalf("could not get user by email: %s", err)
	}
	if getUserRes.ID != createUserRes.ID {
		t.Fatalf("mismatch on ID field; expected '%v', but got '%v'", createUserRes.ID, getUserRes.ID)
	}
	// update user last seen
	lastSeenRes, err := queries.UpdateUserLastSeenAt(context.Background(), getUserRes.ID)
	if err != nil {
		t.Fatalf("could not update user: %s", err)
	}
	if lastSeenRes.LastSeenAt.Before(getUserRes.LastSeenAt) || lastSeenRes.LastSeenAt == getUserRes.LastSeenAt {
		t.Fatal("LastSeenAt field was not successfully updated")
	}
	// this action should also update the updatedAt field
	if lastSeenRes.UpdatedAt.Before(getUserRes.UpdatedAt) || lastSeenRes.UpdatedAt == getUserRes.UpdatedAt {
		t.Fatal("UpdatedAt field was not successfully updated")
	}
}
