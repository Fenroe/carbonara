package config

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/Fenroe/carbonarapi/internal/util"
	"github.com/google/uuid"
)

func (cfg *Config) HandlerCreateUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type responseUser struct {
		ID         uuid.UUID `json:"id"`
		Email      string    `json:"email"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		LastSeenAt time.Time `json:"last_seen_at"`
	}
	type response struct {
		User responseUser `json:"user"`
	}

	body := parameters{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&body)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	hashedPassword, err := auth.HashPassword(body.Password)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't hash password", err)
		return
	}
	queryParams := database.CreateUserParams{
		Email:          body.Email,
		HashedPassword: hashedPassword,
	}
	user, err := cfg.Queries.CreateUser(req.Context(), queryParams)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't create user", err)
		return
	}
	util.RespondWithJSON(w, http.StatusCreated, response{
		User: responseUser{
			ID:         user.ID,
			Email:      user.Email,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			LastSeenAt: user.LastSeenAt,
		},
	})
}
