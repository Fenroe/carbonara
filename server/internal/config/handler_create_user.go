package config

import (
	"encoding/json"
	"net/http"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerCreateUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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
	user, err := cfg.DB.CreateUser(req.Context(), queryParams)
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
