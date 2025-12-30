package main

import (
	"net/http"

	"credilens-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/analyze", handlers.AnalyzeContent)
	}

	r.Run(":8080")

}
