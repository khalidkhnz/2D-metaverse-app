package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	authService "github.com/khalidkhnz/2D-metaverse-app/backend/services/auth"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive" // For ObjectID
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Content-Type for JSON response
		w.Header().Set("Content-Type", "application/json")

		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Authorization token is missing or malformed",
			})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate token
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return lib.GetJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid or expired token",
			})
			return
		}

		// Extract userID from token claims
		userIDStr, ok := (*claims)["userID"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "userID not found in token claims",
			})
			return
		}

		// Convert userID string to ObjectID
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid userID format",
			})
			return
		}

		// Fetch user data from database
		populatedUser, err := authService.GetPopulatedUserByUserId(userID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user", types.FullProfile{
			User:populatedUser,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

