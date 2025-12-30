package models

type AnalyzeRequest struct {
	Content string `json:"content"`
	Type    string `json:"type"` // url | text | image (future)
}

type AnalyzeResponse struct {
	Summary   string   `json:"summary"`
	RiskLevel string   `json:"risk_level"`
	Signals   []string `json:"signals"`
}
