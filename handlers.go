package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chtozamm/rss-aggregator/internal/auth"
	"github.com/chtozamm/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusBadRequest, "Something went wrong")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there!"))
}

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	p := params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse JSON: %s", err))
		return
	}

	insertedUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create new user: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, insertedUser)
}

func (apiCfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, fmt.Sprintf("Authentication error: %s", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("User not found: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
