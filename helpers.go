package main

import (
	"bufio"
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
