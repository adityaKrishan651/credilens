package handlers

import (
	"net/http"

	"credilens-backend/internal/models"
	"credilens-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func AnalyzeContent(c *gin.Context) {
	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	result := services.Analyze(req.Content)

	c.JSON(http.StatusOK, result)
}
