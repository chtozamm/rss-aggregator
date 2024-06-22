package main

import "net/http"

func setupHTTPRouter(ac *apiConfig) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /users", ac.middlewareAuth(ac.getUserHandler))
	r.HandleFunc("POST /users", ac.createUserHandler)
	r.HandleFunc("POST /feeds", ac.middlewareAuth(ac.createFeedHandler))
	r.HandleFunc("GET /feeds", ac.getFeedsHandler)

	return r
}
