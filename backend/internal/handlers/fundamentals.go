package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFundamentals(c *gin.Context) {
	// TODO: fetch fundamentals from Parquet/DB and return JSON
	c.JSON(http.StatusOK, gin.H{"message": "GetFundamentals placeholder"})
}

func ListTickers(c *gin.Context) {
	// TODO: list available tickers
	c.JSON(http.StatusOK, gin.H{"tickers": []string{}})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func ServeSwagger(c *gin.Context) {
	// TODO: serve swagger UI files
	c.JSON(http.StatusOK, gin.H{"swagger": "not implemented"})
}