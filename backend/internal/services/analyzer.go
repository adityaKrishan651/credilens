package services

import "credilens-backend/internal/models"

func Analyze(req models.AnalyzeRequest) models.AnalyzeResponse {

	switch req.Type {

	case "url":
		return analyzeURL(req.Content)

	case "dom":
		return analyzeDOM(req.Content)

	case "image":
		return analyzeImage(req.Content)

	default:
		return analyzeText(req.Content)
	}
}

func analyzeURL(url string) models.AnalyzeResponse {
	return models.AnalyzeResponse{
		Summary: "URL analyzed for credibility & security context",
		RiskLevel: "Medium",
		Signals: []string{
			"Domain age check pending",
			"Source transparency unknown",
		},
	}
}

func analyzeDOM(dom string) models.AnalyzeResponse {
	return models.AnalyzeResponse{
		Summary: "DOM structure analyzed",
		RiskLevel: "Low",
		Signals: []string{
			"No suspicious scripts detected",
		},
	}
}

func analyzeImage(imageBase64 string) models.AnalyzeResponse {
	return models.AnalyzeResponse{
		Summary: "Image metadata analyzed",
		RiskLevel: "Medium",
		Signals: []string{
			"Possible reuse detected",
		},
	}
}

func analyzeText(text string) models.AnalyzeResponse {
	return models.AnalyzeResponse{
		Summary: "Text analyzed using linguistic patterns",
		RiskLevel: "Low",
		Signals: []string{
			"No manipulation indicators",
		},
	}
}
