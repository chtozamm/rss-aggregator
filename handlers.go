package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chtozamm/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (ac *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}

func (ac *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	p := params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to decode request body", err))
		return
	}

	insertedUser, err := ac.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to create new user", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, insertedUser)
}

func (ac *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	p := params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to decode request body", err))
		return
	}

	feed, err := ac.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
		Url:       p.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to create new feed", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

func (ac *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := ac.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to get feeds", err))
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
