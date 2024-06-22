package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// parseEnv reads environmentral variables from the specified file
// and loads them into runtime.
func parseEnv(filename string) error {
	envFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envVar := strings.SplitN(scanner.Text(), "=", 2)
		os.Setenv(purifyString(envVar[0]), purifyString(envVar[1]))
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

// purifyString is a helper function that trims leading and trailing spaces spaces and removes quotation marks from a given string.
func purifyString(s string) string {
	s = strings.Trim(s, " ")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "'", "")
	return s
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

func respondWithError(w http.ResponseWriter, code int, err string) {
	if code > 499 {
		log.Println("Server-side error:", err)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: err,
	})
}

func printErr(message string, err error) string {
	return fmt.Sprint(message, ":", err.Error())
}
