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

func (ac *apiConfig) createFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	p := params{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to decode request body", err))
		return
	}

	feedFollows, err := ac.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    p.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to create new feed follows", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollows)
}

func (ac *apiConfig) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := ac.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to get feeds", err))
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (ac *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIDStr := r.PathValue("id")
	if feedIDStr == "" {
		respondWithError(w, http.StatusBadRequest, printErr("Invalid feed id", nil))
		return
	}

	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to parse feed id", err))
		return
	}

	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	err = ac.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to delete feed", err))
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (ac *apiConfig) getPostsForUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := ac.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, printErr("Failed to get posts", err))
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
