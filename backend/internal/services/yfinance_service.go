package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"yaffa/backend/internal/models"
)

const (
	YFinanceAPIURL = "https://query1.finance.yahoo.com/v8/finance/chart"
)

var yfinanceClient = &http.Client{
	Timeout: 5 * time.Second,
}

// GetYFinanceData fetches real-time market data from yfinance
func GetYFinanceData(ctx context.Context, ticker string) (*models.YFinanceData, error) {
	// TODO: call yfinance API endpoint for ticker
	// TODO: extract price, volume, market cap data
	// TODO: cache results with TTL
	
	url := fmt.Sprintf("%s/%s?interval=1d&range=1y", YFinanceAPIURL, ticker)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := yfinanceClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("yfinance request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("yfinance returned status %d", resp.StatusCode)
	}

	// TODO: parse response and extract latest quote
	yfinanceData := &models.YFinanceData{
		Ticker: ticker,
		Date:   time.Now().Format("2006-01-02"),
		// TODO: populate fields from API response
	}

	return yfinanceData, nil
}

// GetHistoricalPrices fetches historical price data for backtesting
func GetHistoricalPrices(ctx context.Context, ticker string, days int) ([]models.YFinanceData, error) {
	// TODO: implement historical data retrieval
	return nil, nil
}
