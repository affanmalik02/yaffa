package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// LeakyBucket implements a leaky-bucket rate limiter.
type LeakyBucket struct {
	limit      int           // maximum number of requests allowed
	burst      int           // maximum number of requests in a burst
	interval   time.Duration // time interval for the rate limit
	lastUpdate time.Time     // last time the bucket was updated
	tokens     int           // current number of tokens in the bucket
}

// NewLeakyBucket creates a new LeakyBucket instance.
func NewLeakyBucket(limit int, burst int, interval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		limit:    limit,
		burst:    burst,
		interval: interval,
		tokens:   limit,
	}
}

// RateLimitMiddleware enforces the rate limit on incoming requests.
func RateLimitMiddleware(bucket *LeakyBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement leak algorithm and allow/deny requests
		c.Next()
	}
}