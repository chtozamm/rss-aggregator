package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

	respondWithJSON(w, http.StatusOK, insertedUser)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %+v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}
