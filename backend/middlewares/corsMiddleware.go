package middlewares

import (
	"net/http"
)

// CORSOptions holds the configurable options for CORS
type CORSOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// CORSMiddleware creates a CORS middleware for Gorilla Mux router
func CORSMiddleware(options CORSOptions) func(http.Handler) http.Handler {
	// Convert slice of allowed methods to a single comma-separated string
	allowedMethods := "*"
	if len(options.AllowedMethods) > 0 {
		allowedMethods = ""
		for i, method := range options.AllowedMethods {
			if i > 0 {
				allowedMethods += ", "
			}
			allowedMethods += method
		}
	}

	// Convert slice of allowed headers to a single comma-separated string
	allowedHeaders := "*"
	if len(options.AllowedHeaders) > 0 {
		allowedHeaders = ""
		for i, header := range options.AllowedHeaders {
			if i > 0 {
				allowedHeaders += ", "
			}
			allowedHeaders += header
		}
	}

	// Convert slice of allowed origins to a single comma-separated string
	allowedOrigins := "*"
	if len(options.AllowedOrigins) > 0 {
		allowedOrigins = ""
		for i, origin := range options.AllowedOrigins {
			if i > 0 {
				allowedOrigins += ", "
			}
			allowedOrigins += origin
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the CORS headers
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			if options.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Handle preflight (OPTIONS) request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}


