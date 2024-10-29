package config

import (
	"net/http"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerGetClipsByUser(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Clips []responseClip `json:"clips"`
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
	clips, err := cfg.DB.GetClipsByUser(req.Context(), userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't get user clips", err)
		return
	}
	resClips := []responseClip{}
	for _, c := range clips {
		resClips = append(resClips, responseClip{
			ID:        c.ID,
			Content:   c.Content,
			UserID:    c.UserID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	util.RespondWithJSON(w, http.StatusOK, response{
		Clips: resClips,
	})
}
