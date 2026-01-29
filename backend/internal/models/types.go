```go
package models

// Fundamental represents a company's fundamental metrics
type Fundamental struct {
	Ticker          string  `json:"ticker"`
	Date            string  `json:"date"`
	CurrentRatio    float64 `json:"current_ratio"`
	QuickRatio      float64 `json:"quick_ratio"`
	ROE             float64 `json:"roe"`
	NetMargin       float64 `json:"net_margin"`
	PERatio         float64 `json:"pe_ratio"`
	DebtToEquity    float64 `json:"debt_to_equity"`
	FCFYield        float64 `json:"fcf_yield"`
	FairValue       float64 `json:"fair_value"`
}

// YFinanceData represents market data from yfinance
type YFinanceData struct {
	Ticker        string  `json:"ticker"`
	Date          string  `json:"date"`
	Price         float64 `json:"price"`
	Volume        int64   `json:"volume"`
	MarketCap     float64 `json:"market_cap"`
	Dividend      float64 `json:"dividend"`
}

// PlaidHolding represents user portfolio holding synced via Plaid
type PlaidHolding struct {
	UserID        string  `json:"user_id"`
	SecurityID    string  `json:"security_id"`
	Ticker        string  `json:"ticker"`
	Quantity      float64 `json:"quantity"`
	CostBasis     float64 `json:"cost_basis"`
	CurrentValue  float64 `json:"current_value"`
}
```