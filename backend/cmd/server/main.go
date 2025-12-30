package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"credilens-backend/internal/clients"
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

	openAIClient, err := clients.NewOpenAIClient(
		cfg.AzureOpenAIEndpoint,
		cfg.AzureOpenAIKey,
		cfg.AzureOpenAIDeploy,
	)

	if err != nil {
		log.Println("⚠️ Azure OpenAI not configured, AI features disabled")
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	{
		var aiChatFunc handlers.AIChatFunc

		if openAIClient != nil {
			aiChatFunc = openAIClient.Chat
		}

		api.POST("/analyze", handlers.AnalyzeContent(aiChatFunc))
	}

	r.Run(":" + cfg.ServerPort)
}
