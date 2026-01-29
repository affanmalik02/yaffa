package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIKeyMiddleware is a middleware function that checks for a valid API key in the request headers.
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the API key from the request header
		apiKey := c.Request.Header.Get("X-API-KEY")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing X-API-KEY header"})
			c.Abort()
			return
		}

		// Validate the API key
		isValid, err := ValidateAPIKey(apiKey)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			c.Abort()
			return
		}

		// If the API key is valid, proceed to the next middleware or handler
		c.Next()
	}
}

// ValidateAPIKey checks the validity of the given API key.
// It queries the MongoDB to verify if the API key exists and is active.
func ValidateAPIKey(key string) (bool, error) {
	// TODO: Implement MongoDB query to validate API key
	return false, nil
}