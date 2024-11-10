package lib

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
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
		"exp":    GetExpirationTime(), 
	}

	// Create the JWT token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}


func UserInContext(w http.ResponseWriter, r *http.Request) (*types.FullProfile) {
	user, ok := r.Context().Value("user").(types.FullProfile)
	if !ok {
		WriteJSON(w, http.StatusUnauthorized, "User not found in context")
		return &types.FullProfile{}
	}
	return &user
}


func GenerateShortLivedJwtToken(user *types.FullProfile) (string, error) {
	// Create token with claims
	claims := jwt.MapClaims{
		"userID": user.User.ID,
		"isWS": true,
		"exp":    GetShortLivedExpirationTime(), 
	}

	// Create the JWT token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}