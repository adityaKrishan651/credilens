package models

type AnalyzeRequest struct {
	Type    string `json:"type" binding:"required,oneof=url dom image text"`
	Content string `json:"content" binding:"required"`
}

type AnalyzeResponse struct {
	Summary   string   `json:"summary"`
	RiskLevel string   `json:"risk_level"`
	Signals   []string `json:"signals"`
}
