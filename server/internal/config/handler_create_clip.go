package config

import (
	"encoding/json"
	"net/http"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerCreateClip(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Content string `json:"content"`
	}

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
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	clip, err := cfg.DB.CreateClip(req.Context(), database.CreateClipParams{
		Content: params.Content,
		UserID:  userID,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't create clip", err)
		return
	}
	util.RespondWithJSON(w, http.StatusCreated, response{
		Clip: responseClip{
			ID:        clip.ID,
			Content:   clip.Content,
			UserID:    clip.UserID,
			CreatedAt: clip.CreatedAt,
			UpdatedAt: clip.UpdatedAt,
		},
	})
}
