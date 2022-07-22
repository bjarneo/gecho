package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define the middlewares
	r.Use(LoggingMiddleware)

	r.Handle("/healthz", http.HandlerFunc(Healthz))

	// Handle all routes, no matter what
	r.PathPrefix("/").Handler(http.HandlerFunc(Echo))

	http.ListenAndServe(":"+Port(), handlers.CompressHandler(r))
}
