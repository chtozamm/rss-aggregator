package main

import (
	"net/http"

	"github.com/chtozamm/rss-aggregator/internal/auth"
	"github.com/chtozamm/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ac *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, printErr("Failed to get api key", err))
			return
		}

		user, err := ac.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, printErr("Failed to get user by api key", err))
			return
		}

		handler(w, r, user)
	}
}
