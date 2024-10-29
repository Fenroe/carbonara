package config

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerLogin(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		User         responseUser `json:"user"`
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "coulkdn't decode parameters", err)
		return
	}
	user, err := cfg.DB.GetUserByEmail(req.Context(), params.Email)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "incorrect email or password", err)
		return
	}
	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "incorrect email or password", err)
		return
	}
	accessToken, err := auth.MakeJWT(
		user.ID,
		cfg.JWTSecret,
		time.Hour,
	)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't create access token", err)
		return
	}
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't create refresh token", err)
		return
	}
	_, err = cfg.DB.CreateRefreshToken(req.Context(), database.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't save refresh token", err)
		return
	}
	// Everything else has succeeded so update user last seen date
	user, err = cfg.DB.UpdateUserLastSeenAt(req.Context(), user.ID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't update user data", err)
	}
	util.RespondWithJSON(w, http.StatusOK, response{
		User: responseUser{
			ID:         user.ID,
			Email:      user.Email,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			LastSeenAt: user.LastSeenAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
