package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chtozamm/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environmental variables
	if err := parseEnv(".env"); err != nil {
		log.Fatalf(printErr("Failed to parse .env file", err))
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
		log.Fatal(printErr("Failed to connect to the database", err))
	}

	db := database.New(conn)
	ac := apiConfig{DB: db}

	// Setup http router
	r := setupHTTPRouter(&ac)

	go startScraping(db, 10, time.Minute)

	log.Println("Server is listening on http://localhost:" + port)
	err = http.ListenAndServe("localhost:"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
