package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

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


func RunReactApp() error {
	// Change directory to ./views
	if err := os.Chdir("./views"); err != nil {
		log.Printf("Error changing directory to './views': %v", err)
		return err
	}
	log.Printf("Changed directory to './views'")

	// Run the React app
	cmd := exec.Command("sh", "-c", "npm start")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running 'npm start' command: %v", err)
		return err
	}
	log.Printf("Output of 'npm start' command: %s", string(output))

	return nil
}
