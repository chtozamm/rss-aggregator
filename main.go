package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environmental variables
	if err := parseEnv(".env"); err != nil {
		log.Fatalf("Failed to parse env file: %s", err)
	}

	// Setup http router
	r := http.NewServeMux()
	r.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there!"))
	})
	r.HandleFunc("GET /health", checkHealthHandler)
	r.HandleFunc("GET /error", errorHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is listening on http://localhost:" + port)
	http.ListenAndServe("localhost:"+port, r)
}
