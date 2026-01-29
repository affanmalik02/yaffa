package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CacheKey generates a cache key for a given entity
func CacheKey(entity, identifier string) string {
	return fmt.Sprintf("%s:%s", entity, identifier)
}

// GetCachedFundamentals retrieves fundamentals from cache (MongoDB)
func GetCachedFundamentals(ctx context.Context, ticker string) (map[string]interface{}, error) {
	// TODO: query MongoDB cache collection with key
	// TODO: check TTL and return if not expired
	return nil, nil
}

// SetCachedFundamentals stores fundamentals in cache with TTL
func SetCachedFundamentals(ctx context.Context, ticker string, data map[string]interface{}, ttl time.Duration) error {
	// TODO: store in MongoDB with expiration time
	return nil
}

// GetCachedPrice retrieves cached yfinance price data
func GetCachedPrice(ctx context.Context, ticker string) (float64, error) {
	// TODO: query MongoDB price cache
	return 0, nil
}

// SetCachedPrice stores price data with short TTL (5-15 minutes)
func SetCachedPrice(ctx context.Context, ticker string, price float64, ttl time.Duration) error {
	// TODO: store price with TTL
	return nil
}
