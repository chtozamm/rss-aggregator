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
		envVar := strings.Split(scanner.Text(), "=")
		os.Setenv(strings.Trim(envVar[0], " "), strings.Trim(envVar[1], " "))
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}
