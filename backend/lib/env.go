package lib

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// InitEnv loads environment variables from .env file
func InitEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

var (
    EXPIRATION_DURATION = time.Hour * 24 * 7
    SHORT_EXPIRATION_DURATION = time.Hour / 4
)

// Functions to retrieve environment variables after InitEnv() is called
func GetJWTSecret() []byte {
    return []byte(os.Getenv("JWT_SECRET"))
}

func GetPort() string {
    return ":" + getEnv("PORT", "4000")
}

func GetAPIPrefix() string {
    return getEnv("API_PREFIX", "/api/v1")
}

func GetAPIBase() string {
    return getEnv("API_BASE_URL", "")
}

func GetFullAPIBase() string {
    return getEnv("FULL_API_BASE_URL", "")
}

func GetFrontendProxyURL() string {
    return getEnv("FRONTEND_PROXY_URL", "")
}

func GetDBURI() string {
    return getEnv("DB_URI", "")
}

func GetExpirationTime() int64 {
    return time.Now().Add(EXPIRATION_DURATION).Unix()
}

func GetShortLivedExpirationTime() int64 {
    return time.Now().Add(SHORT_EXPIRATION_DURATION).Unix()
}

// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key string, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
