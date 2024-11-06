package test

import (
	"context"
	"testing"
	"time"

	"github.com/Fenroe/carbonarapi/internal/database"
)

func TestRefreshTokens(t *testing.T) {
	// setup and cleanup
	queries, cleanup := dbTestSetup()
	defer cleanup()
	// insert user as refresh token model depends on it
	user, err := queries.CreateUser(context.Background(), database.CreateUserParams{
		Email:          "test@carbonara.com",
		HashedPassword: "myhashedpassword",
	})
	if err != nil {
		t.Fatalf("could not create test dependencies: %s", err)
	}
	// create refresh token
	createTokenRes, err := queries.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     "refreshtokenstring",
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		t.Fatalf("could not create refresh token: %s", err)
	}
	// token field should be unique
	_, err = queries.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     createTokenRes.Token,
		UserID:    createTokenRes.UserID,
		ExpiresAt: createTokenRes.ExpiresAt,
	})
	if err == nil {
		t.Fatalf("refresh token with duplicate token field was created")
	}
	// the userID should not be unique
	_, err = queries.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     "anewtokenstring",
		UserID:    createTokenRes.UserID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		t.Fatalf("refresh token with matching userID to another token could not be created")
	}
	// extend refresh token
	extendedRefreshToken, err := queries.ExtendRefreshToken(context.Background(), database.ExtendRefreshTokenParams{
		ExpiresAt: createTokenRes.ExpiresAt.Add(time.Hour * 24 * 30),
		Token:     createTokenRes.Token,
	})
	if err != nil {
		t.Fatalf("could not extend refresh token: %s", err)
	}
	if extendedRefreshToken.ExpiresAt.Before(createTokenRes.ExpiresAt) || extendedRefreshToken.ExpiresAt == createTokenRes.ExpiresAt {
		t.Fatal("expected ExpiresAt field on the token to be greater than it was previously, but it is not")
	}
}
