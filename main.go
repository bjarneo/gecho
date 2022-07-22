package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Fetch the port from env, if not set, return a default port
func port() string {
	port := os.Getenv("HTTP_PORT")

	if port != "" {
		return port
	}

	return "8080"
}

// Sort the header keys
func sortKeyHeaders(headers http.Header) []string {
	keys := make([]string, 0, len(headers))

	for name := range headers {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	return keys
}

// The echo endpoint
func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "=========================\n")
	fmt.Fprintf(w, "|--- REQUEST DETAILS ---|\n")
	fmt.Fprintf(w, "=========================\n")

	fmt.Fprintf(w, "URI: %s", req.URL.RequestURI())

	fmt.Fprintf(w, "\n\n=========================\n")
	fmt.Fprintf(w, "|--- REQUEST HEADERS ---|\n")
	fmt.Fprintf(w, "=========================\n")

	// Sort the header keys
	sorted := sortKeyHeaders(req.Header)
	headers := req.Header

	for _, key := range sorted {
		// Iterate through the sorted keys,
		// then use the key in the headers range
		// to get alphabetically order
		for _, h := range headers[key] {
			fmt.Fprintf(w, "%v: %v\n", key, h)
		}
	}

	body, _ := ioutil.ReadAll(req.Body)

	// If we have a request body, print it
	if len(body) > 0 {
		fmt.Fprintf(w, "\n\n=========================\n")
		fmt.Fprintf(w, "|---- REQUEST  BODY ----|\n")
		fmt.Fprintf(w, "=========================\n")
		fmt.Fprintf(w, "%s", body)
	}
}

// Just a basic health endpoint
func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I am OK 🥲\n")
}

func main() {
	r := mux.NewRouter()

	r.Handle("/healthz", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(healthz)))

	// Handle all routes, no matter what
	r.PathPrefix("/").Handler(handlers.LoggingHandler(os.Stdout, http.HandlerFunc(echo)))

	http.ListenAndServe(":"+port(), handlers.CompressHandler(r))
}
