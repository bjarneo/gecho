package main

import (
	"net/http"
	"os"
	"sort"
)

// Fetch the port from env, if not set, return a default port
func Port() string {
	port := os.Getenv("HTTP_PORT")

	if port != "" {
		return port
	}

	return "8080"
}

// Sort the header keys
func SortKeyHeaders(headers http.Header) []string {
	keys := make([]string, 0, len(headers))

	for name := range headers {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	return keys
}
