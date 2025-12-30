package models

type AnalyzeRequest struct {
	Content string `json:"content"`
	Type    string `json:"type"` // url | text | image (future)
	Source  string `json:"source"`
}

type AnalyzeResponse struct {
	Summary   string    `json:"summary"`
	RiskLevel RiskLevel `json:"risk_level"`
	Signals   []Signal  `json:"signals"`
}
