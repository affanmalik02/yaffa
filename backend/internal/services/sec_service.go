package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"yaffa/backend/internal/models"
)

const (
	SECBaseURL = "https://data.sec.gov/submissions"
	SECUserAgent = "YAFFA (Affan Malik, affan.malik@example.com)"
)

var secClient = &http.Client{
	Timeout: 10 * time.Second,
}

// GetSECFundamentals fetches company fundamentals from SEC EDGAR
func GetSECFundamentals(ctx context.Context, ticker string) (*models.Fundamental, error) {
	// TODO: implement CIK lookup from ticker
	// TODO: fetch company facts JSON from SEC API
	// TODO: parse XBRL tags and extract financial metrics
	// TODO: cache results in MongoDB
	
	cik, err := getTicKertoCIK(ctx, ticker)
	if err != nil {
		return nil, fmt.Errorf("failed to map ticker to CIK: %w", err)
	}

	facts, err := fetchCompanyFacts(ctx, cik)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company facts: %w", err)
	}

	fundamental := parseXBRLFacts(facts)
	return fundamental, nil
}

// getTicKertoCIK maps ticker to SEC CIK identifier
func getTicKertoCIK(ctx context.Context, ticker string) (string, error) {
	// TODO: query mapping table or cache
	// For now, placeholder
	return "", nil
}

// fetchCompanyFacts retrieves raw company facts from SEC API
func fetchCompanyFacts(ctx context.Context, cik string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/CIK%s/company_facts.json", SECBaseURL, padCIK(cik))
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", SECUserAgent)

	resp, err := secClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var facts map[string]interface{}
	if err := json.Unmarshal(body, &facts); err != nil {
		return nil, err
	}

	return facts, nil
}

// parseXBRLFacts converts raw XBRL JSON into Fundamental struct
func parseXBRLFacts(facts map[string]interface{}) *models.Fundamental {
	// TODO: implement XBRL tag parsing
	// TODO: extract NetIncome, Assets, Liabilities, etc.
	return &models.Fundamental{
		// TODO: populate fields
	}
}

// padCIK pads CIK to 10-digit format required by SEC API
func padCIK(cik string) string {
	// TODO: implement padding logic
	return cik
}
