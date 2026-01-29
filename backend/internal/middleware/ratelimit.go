package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type LeakyBucket struct {
	limit     int
	burst     int
	tokens    int
	mu        sync.Mutex
	lastDrain time.Time
}

// NewLeakyBucket creates a new LeakyBucket instance.
func NewLeakyBucket(limit int, burst int) *LeakyBucket {
	return &LeakyBucket{
		limit:     limit,
		burst:     burst,
		tokens:    burst,
		lastDrain: time.Now(),
	}
}

// RateLimitMiddleware enforces the rate limit on incoming requests.
func RateLimitMiddleware(bucket *LeakyBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		bucket.mu.Lock()
		defer bucket.mu.Unlock()

		now := time.Now()
		elapsed := now.Sub(bucket.lastDrain).Seconds()
		bucket.tokens = min(bucket.burst, bucket.tokens+int(elapsed*float64(bucket.limit)))
		bucket.lastDrain = now

		if bucket.tokens > 0 {
			bucket.tokens--
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}