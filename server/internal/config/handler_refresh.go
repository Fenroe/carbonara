package config

import (
	"net/http"
	"time"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerRefresh(w http.ResponseWriter, req *http.Request) {
	type response struct {
		AccessToken string `json:"access_token"`
	}

	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "couldn't get refresh token", err)
		return
	}
	user, err := cfg.DB.GetUserFromRefreshToken(req.Context(), refreshToken)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "couldn't get user for refresh token", err)
		return
	}
	accessToken, err := auth.MakeJWT(
		user.ID,
		cfg.JWTSecret,
		time.Hour,
	)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "couldn't validate token", err)
		return
	}
	// Extend the refresh token as it's still in use
	_, err = cfg.DB.ExtendRefreshToken(req.Context(), database.ExtendRefreshTokenParams{
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 30),
		Token:     refreshToken,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't extend refresh token", err)
		return
	}
	util.RespondWithJSON(w, http.StatusOK, response{
		AccessToken: accessToken,
	})
}
