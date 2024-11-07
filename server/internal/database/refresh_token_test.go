package database

import (
	"context"
	"testing"
	"time"
)

func TestRefreshTokens(t *testing.T) {
	// setup and cleanup
	queries, cleanup := DBTestSetup()
	defer cleanup()
	// insert user as refresh token model depends on it
	user, err := queries.CreateUser(context.Background(), CreateUserParams{
		Email:          "test@carbonara.com",
		HashedPassword: "myhashedpassword",
	})
	if err != nil {
		t.Fatalf("could not create test dependencies: %s", err)
	}
	// create refresh token
	createTokenRes, err := queries.CreateRefreshToken(context.Background(), CreateRefreshTokenParams{
		Token:     "refreshtokenstring",
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		t.Fatalf("could not create refresh token: %s", err)
	}
	// token field should be unique
	_, err = queries.CreateRefreshToken(context.Background(), CreateRefreshTokenParams{
		Token:     createTokenRes.Token,
		UserID:    createTokenRes.UserID,
		ExpiresAt: createTokenRes.ExpiresAt,
	})
	if err == nil {
		t.Fatalf("refresh token with duplicate token field was created")
	}
	// the userID should not be unique
	_, err = queries.CreateRefreshToken(context.Background(), CreateRefreshTokenParams{
		Token:     "anewtokenstring",
		UserID:    createTokenRes.UserID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		t.Fatalf("refresh token with matching userID to another token could not be created")
	}
	// extend refresh token
	extendedRefreshToken, err := queries.ExtendRefreshToken(context.Background(), ExtendRefreshTokenParams{
		ExpiresAt: createTokenRes.ExpiresAt.Add(time.Hour * 24 * 30),
		Token:     createTokenRes.Token,
	})
	if err != nil {
		t.Fatalf("could not extend refresh token: %s", err)
	}
	if extendedRefreshToken.ExpiresAt.Before(createTokenRes.ExpiresAt) || extendedRefreshToken.ExpiresAt == createTokenRes.ExpiresAt {
		t.Fatal("expected ExpiresAt field on the token to be greater than it was previously, but it is not")
	}
	// find user by refresh token
	foundUser, err := queries.GetUserFromRefreshToken(context.Background(), extendedRefreshToken.Token)
	if err != nil {
		t.Fatalf("could not find user with refresh token: %s", err)
	}
	if foundUser.ID != user.ID {
		t.Fatal("could not match the expected user with refresh token")
	}
	// revoke refresh token
	_, err = queries.RevokeRefreshToken(context.Background(), extendedRefreshToken.Token)
	if err != nil {
		t.Fatalf("could not revoke refresh token: %s", err)
	}
	// revoked refresh token cannot be used for other queries
	_, err = queries.GetUserFromRefreshToken(context.Background(), extendedRefreshToken.Token)
	if err == nil {
		t.Fatal("revoked refresh token was used to find user")
	}
	_, err = queries.ExtendRefreshToken(context.Background(), ExtendRefreshTokenParams{
		ExpiresAt: createTokenRes.ExpiresAt.Add(time.Hour * 24 * 30),
		Token:     createTokenRes.Token,
	})
	if err == nil {
		t.Fatal("revoked refresh token was extended")
	}
}
