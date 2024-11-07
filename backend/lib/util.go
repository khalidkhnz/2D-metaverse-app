package lib

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}



func GenerateJWTToken(userID string) (string, error) {
	// Create token with claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    ExpirationTime, 
	}

	// Create the JWT token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

