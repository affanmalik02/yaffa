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

		// Validate the API key
		valid, err := ValidateAPIKey(apiKey)
		if err != nil || !valid {
			// If the API key is invalid or there is an error, abort the request with a 401 Unauthorized status
			c.AbortWithStatus(http.StatusUnauthorized)
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