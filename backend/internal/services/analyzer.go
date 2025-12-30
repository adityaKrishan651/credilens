package services

import (
	"fmt"

	"credilens-backend/internal/models"
	"credilens-backend/internal/prompts"
)

func AnalyzeText(
	openAI func(system, user string) (string, error),
	content string,
	source string,
) models.AnalyzeResponse {

	if openAI == nil {
		return models.AnalyzeResponse{
			RiskLevel: models.RiskMedium,
			Summary:   "AI analysis is currently unavailable.",
			Signals: []models.Signal{
				{
					Code:        "AI_DISABLED",
					Description: "AI analysis is not configured.",
					Severity:    "medium",
					Source:      "system",
				},
			},
		}
	}
	aiOutput, err := openAI(
		prompts.CredibilitySystemPrompt, fmt.Sprintf(prompts.CredibilitySystemPrompt, content),
	)

	signals := []models.Signal{}

	if err != nil {
		signals = append(signals, models.Signal{
			Code:        "AI_UNAVAILABLE",
			Description: "AI analysis coudl not be performed.",
			Severity:    "medium",
			Source:      "system",
		})

		return models.AnalyzeResponse{
			RiskLevel: models.RiskMedium,
			Summary:   "AI Analysis unavailable. Showing limited signals.",
			Signals:   signals,
		}
	}

	signals = append(signals, models.Signal{
		Code:        "AI_ANALYSIS_COMPLETE",
		Description: "AI analysis completed successfully.",
		Severity:    "low",
		Source:      "ai",
	})

	return models.AnalyzeResponse{
		RiskLevel: models.RiskMedium,
		Summary:   aiOutput,
		Signals:   signals,
	}
}

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
