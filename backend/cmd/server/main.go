package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"credilens-backend/internal/config"
	"credilens-backend/internal/handlers"
)

func main() {
	cfg := config.Load()
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/analyze", handlers.AnalyzeContent)
	}

	r.Run(":" + cfg.ServerPort)
}
