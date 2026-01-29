package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaffa/backend/internal/models"
	"yaffa/backend/internal/services"
)

func GetFundamentals(c *gin.Context) {
	ticker := c.Param("ticker")
	
	// TODO: fetch fundamentals from S3 Parquet store
	// TODO: fetch real-time price from yfinance
	// TODO: merge SEC fundamentals with market data
	secData, err := services.GetSECFundamentals(c.Request.Context(), ticker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "fundamentals not found"})
		return
	}

	priceData, err := services.GetYFinanceData(c.Request.Context(), ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch price data"})
		return
	}

	merged := mergeData(secData, priceData)
	c.JSON(http.StatusOK, merged)
}

func ListTickers(c *gin.Context) {
	// TODO: list available tickers from metadata collection
	tickers, err := services.GetAvailableTickers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tickers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tickers": tickers})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "yaffa-api"})
}

func ServeSwagger(c *gin.Context) {
	// TODO: serve swagger UI files from static directory
	c.JSON(http.StatusOK, gin.H{"swagger": "available at /swagger/index.html"})
}

// mergeData combines SEC fundamentals with yfinance market data
func mergeData(sec *models.Fundamental, market *models.YFinanceData) map[string]interface{} {
	return map[string]interface{}{
		"ticker":          sec.Ticker,
		"date":            sec.Date,
		"fundamentals":    sec,
		"market_data":     market,
		"fair_value":      sec.FairValue,
		"current_price":   market.Price,
		"valuation_gap":   sec.FairValue - market.Price,
	}
}