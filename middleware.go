package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

var maxFileSize int64 = 1024 * 1024 * 2 // 2MB

func MaxFileSize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
		next.ServeHTTP(w, r)
	})
}
