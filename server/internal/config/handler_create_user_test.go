package config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestHandlerCreateUser(t *testing.T) {
	queries, cleanup := database.DBTestSetup()
	defer cleanup()
	cfg := &Config{DB: queries}

	// Define request parameters
	requestBody := map[string]string{
		"email":    "test@example.com",
		"password": "securepassword",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Define the expected user to be returned by the mock database
	expectedUser := database.User{
		Email:      "test@example.com",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		LastSeenAt: time.Now(),
	}

	// Create request and recorder
	req := httptest.NewRequest(http.MethodPost, "/create-user", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Call the handler
	cfg.HandlerCreateUser(recorder, req)

	// Check response
	res := recorder.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// Parse the response JSON
	var response struct {
		User responseUser `json:"user"`
	}
	err := json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	// Validate the response user data
	assert.Equal(t, expectedUser.Email, response.User.Email)
	assert.WithinDuration(t, expectedUser.CreatedAt, response.User.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedUser.UpdatedAt, response.User.UpdatedAt, time.Second)
	assert.WithinDuration(t, expectedUser.LastSeenAt, response.User.LastSeenAt, time.Second)
}
