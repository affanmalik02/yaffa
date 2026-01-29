```go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"yaffa/backend/internal/db"
	"yaffa/backend/internal/handlers"
	"yaffa/backend/internal/middleware"
)

// @title YAFFA API
// @version 1.0
// @description This is the API documentation for the YAFFA project.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.yourdomain.com/support
// @contact.email support@yourdomain.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Initialize MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := db.Connect(ctx, mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Initialize router
	r := setupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting YAFFA API server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.APIKeyMiddleware())
	bucket := middleware.NewLeakyBucket(10, 20) // 10 req/s with burst of 20
	r.Use(middleware.RateLimitMiddleware(bucket))

	// Routes
	registerRoutes(r)

	return r
}

func registerRoutes(r *gin.Engine) {
	// Health check (no auth required)
	r.GET("/health", handlers.HealthCheck)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// Fundamentals endpoints
		v1.GET("/fundamentals/:ticker", handlers.GetFundamentals)
		v1.GET("/tickers", handlers.ListTickers)

		// Swagger documentation
		v1.GET("/swagger", handlers.ServeSwagger)
	}
}
```