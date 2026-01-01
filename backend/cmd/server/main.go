package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"credilens-backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Fail fast if Azure keys are missing
	requiredEnvs := []string{
		"AZURE_OPENAI_KEY",
		"AZURE_AI_SEARCH_KEY",
		"AZURE_AI_VISION_KEY",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("Missing required env var: %s", env)
		}
	}

	r := gin.Default()

	// CORS (Browser extension + Web app safe)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"chrome-extension://*",
		},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		MaxAge: 12 * time.Hour,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/analyze", handlers.AnalyzeContent)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("CrediLens backend running on port", port)
	r.Run(":" + port)
}
