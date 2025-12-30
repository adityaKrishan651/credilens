package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"credilens-backend/internal/models"
	"credilens-backend/internal/services"
)

type AIChatFunc func(system, user string) (string, error)

func AnalyzeContent(aiChat AIChatFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AnalyzeRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Invalid request payload",
			})
			return
		}

		// validation already exists
		if err := services.ValidateAnalyzeInput(req.Type, req.Content); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		var resp models.AnalyzeResponse

		if req.Type == "text" || req.Type == "dom" {
			resp = services.AnalyzeText(aiChat, req.Content, req.Source)
		} else {
			resp = services.Analyze(req.Type, req.Content, req.Source)
		}

		c.JSON(http.StatusOK, resp)
	}
}
