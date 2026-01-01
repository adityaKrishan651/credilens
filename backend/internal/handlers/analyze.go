package handlers

import (
	"net/http"
	"net/url"

	"credilens-backend/internal/models"
	"credilens-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func AnalyzeContent(c *gin.Context) {
	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Type-specific validation
	if req.Type == "url" {
		if _, err := url.ParseRequestURI(req.Content); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
	}

	result := services.Analyze(req)

	c.JSON(http.StatusOK, result)
}
