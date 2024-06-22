package main

import "net/http"

func setupHTTPRouter(ac *apiConfig) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /{$}", homeHandler)
	r.HandleFunc("GET /health", checkHealthHandler)
	r.HandleFunc("GET /error", errorHandler)
	r.HandleFunc("POST /users", ac.createUserHandler)

	return r
}
