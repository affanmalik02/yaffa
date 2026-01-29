```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	r := setupRouter()
	// TODO: configure port from env/config
	_ = r.Run() // default :8080
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	registerRoutes(r)
	return r
}

func registerRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// TODO: register other routes
}
```