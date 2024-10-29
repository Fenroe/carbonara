package config

import (
	"net/http"

	"github.com/Fenroe/carbonarapi/internal/auth"
	"github.com/Fenroe/carbonarapi/internal/util"
)

func (cfg *Config) HandlerRevoke(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "couldn't find token", err)
		return
	}
	_, err = cfg.DB.RevokeRefreshToken(req.Context(), refreshToken)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "couldn't revoke session", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
