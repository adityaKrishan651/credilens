package handlers

import (
	"net/http"

	"credilens-backend/internal/models"
	"credilens-backend/internal/services"
	"credilens-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func AnalyzeContent(c *gin.Context) {
	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Inavlid request payload",
		})
		return
	}

	if err := services.ValidateAnalyseInput(req.Type, req.Content); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	contentHash := utils.HashContent(req.Content)
	_ = contentHash

	resp := services.Analyze(req.Type, req.Content, req.Source)
	c.JSON(http.StatusOK, resp)
}
