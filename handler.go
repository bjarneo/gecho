package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// The echo endpoint
func Echo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "=========================\n")
	fmt.Fprintf(w, "|--- REQUEST DETAILS ---|\n")
	fmt.Fprintf(w, "=========================\n")

	fmt.Fprintf(w, "URI: %s", req.URL.RequestURI())

	fmt.Fprintf(w, "\n\n=========================\n")
	fmt.Fprintf(w, "|--- REQUEST HEADERS ---|\n")
	fmt.Fprintf(w, "=========================\n")

	// Sort the header keys
	sorted := SortKeyHeaders(req.Header)
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
func Healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I am OK 🥲\n")
}
