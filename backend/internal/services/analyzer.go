package services

import "credilens-backend/internal/models"

func Analyze(inputType string, content string, source string) models.AnalyzeResponse {
	signals := []models.Signal{
		{
			Code:        "CONTENT_RECEIVED",
			Description: "The content was successfully received and validated.",
			Severity:    "low",
			Source:      "system",
		},
		{
			Code:        "INPUT_TYPE_" + inputType,
			Description: "Content type identified as " + inputType + ".",
			Severity:    "low",
			Source:      "system",
		},
	}

	if source != "" {
		signals = append(signals, models.Signal{
			Code:        "SOURCE_" + source,
			Description: "Content submitted via " + source + ".",
			Severity:    "low",
			Source:      "system",
		})
	}

	return models.AnalyzeResponse{
		RiskLevel: models.RiskLow,
		Summary:   "Content has passed initial validation and is ready for analysis.",
		Signals:   signals,
	}
}
