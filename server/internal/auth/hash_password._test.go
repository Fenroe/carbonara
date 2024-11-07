package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordTest(t *testing.T) {
	// test data
	password1 := "password123!"

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "Password too long",
			password: password1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if hash == tt.password {
				t.Error("CheckPasswordHash() hash and password match")
			}
			if hash == "" {
				t.Error("CheckPasswordHash() hash is empty")
			}
			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.password))
			if err != nil {
				t.Error("CheckPasswordHash() does not produce a hashed password")
			}
		})
	}
}
