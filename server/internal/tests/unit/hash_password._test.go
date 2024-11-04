package main

import (
	"testing"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPasswordSuccess(t *testing.T) {
	password := "securePassword123"
	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if hash == "" {
		t.Fatal("Expected non-empty hash, but got empty string")
	}

	// Verify the hash is valid by using bcrypt's CompareHashAndPassword
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Fatalf("Expected hash to match password, but got error: %v", err)
	}
}
