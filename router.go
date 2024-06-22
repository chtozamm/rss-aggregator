package main

import "net/http"

func setupHTTPRouter(ac *apiConfig) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /users", ac.middlewareAuth(ac.getUserHandler))
	r.HandleFunc("POST /users", ac.createUserHandler)
	r.HandleFunc("GET /feeds", ac.getFeedsHandler)
	r.HandleFunc("POST /feeds", ac.middlewareAuth(ac.createFeedHandler))
	r.HandleFunc("GET /feed_follows", ac.middlewareAuth(ac.getFeedFollowsHandler))
	r.HandleFunc("POST /feed_follows", ac.middlewareAuth(ac.createFeedFollowsHandler))
	r.HandleFunc("DELETE /feed_follows/{id}", ac.middlewareAuth(ac.deleteFeedFollowHandler))

	return r
}
