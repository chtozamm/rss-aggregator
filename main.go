package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/chtozamm/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environmental variables
	if err := parseEnv(".env"); err != nil {
		log.Fatalf("Failed to parse env file: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	apiCfg := apiConfig{DB: database.New(conn)}

	// Setup http router
	r := setupHTTPRouter(&apiCfg)

	log.Println("Server is listening on http://localhost:" + port)
	http.ListenAndServe("localhost:"+port, r)
}
