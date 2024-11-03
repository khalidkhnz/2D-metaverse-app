package middlewares

import (
	"log"
	"net/http"
	"time"
)

// ANSI escape codes for coloring
const (
	greenBackground  = "\033[42m" // Green background
	whiteText        = "\033[37m" // White text
	yellowText       = "\033[33m" // Yellow text for status code
	cyanText         = "\033[36m" // Cyan text for processing time
	reset            = "\033[0m"   // Reset color
)

// LoggingMiddleware logs each API call with its URL, status, method, and processing time
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a ResponseRecorder to capture the status code
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler in the chain
		next.ServeHTTP(rec, r)

		// Log the request details with colored method, status code, and time
		log.Printf(
			"%s%s%s %s%d%s %s%s in %v%s",
			greenBackground, r.Method, reset,          // Green background for method
			yellowText, rec.statusCode, reset,         // Yellow text for status code
			r.RemoteAddr,                               // Remote address
			cyanText, time.Since(start), reset,        // Cyan text for processing time
		)
	})
}

// responseRecorder is a custom ResponseWriter that captures the status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
