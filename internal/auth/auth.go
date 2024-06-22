package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts API key from the headers of an HTTP request.
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing \"Authorization\" header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed \"Authorization\" header")
	}

	if strings.ToLower(vals[0]) != "apikey" {
		return "", errors.New("malformed \"Authorization\" header")
	}

	return vals[1], nil
}
