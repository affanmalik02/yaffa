package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"yaffa/backend/internal/models"
)

var plaidClient = &http.Client{
	Timeout: 10 * time.Second,
}

// ExchangePlaidToken exchanges Plaid public token for access token
func ExchangePlaidToken(ctx context.Context, publicToken string) (string, error) {
	// TODO: call Plaid token exchange endpoint
	// TODO: store encrypted access token in MongoDB
	return "", nil
}

// GetPlaidHoldings retrieves portfolio holdings via Plaid
func GetPlaidHoldings(ctx context.Context, userID string) ([]models.PlaidHolding, error) {
	// TODO: fetch stored Plaid access token from MongoDB
	// TODO: call Plaid holdings endpoint
	// TODO: parse response and map security IDs to tickers
	return nil, nil
}

// SyncPlaidAccount performs full account sync for user
func SyncPlaidAccount(ctx context.Context, userID string) error {
	holdings, err := GetPlaidHoldings(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch Plaid holdings: %w", err)
	}

	// TODO: map each holding to ticker via security ID
	// TODO: fetch current prices from yfinance
	// TODO: store in MongoDB portfolio collection

	for _, holding := range holdings {
		_ = holding // TODO: process each holding
	}

	return nil
}
