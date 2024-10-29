package config

import (
	"net/http"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerGetMostRecentClip(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Clip responseClip `json:"clip"`
	}
	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.JWTSecret)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "couldn't validate JWT", err)
		return
	}
	clip, err := cfg.DB.GetMostRecentClip(req.Context(), userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't get user clips", err)
		return
	}
	util.RespondWithJSON(w, http.StatusOK, response{
		Clip: responseClip{
			ID:        clip.ID,
			Content:   clip.Content,
			UserID:    clip.UserID,
			CreatedAt: clip.CreatedAt,
			UpdatedAt: clip.UpdatedAt,
		},
	})
}
