package models

type Signal struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Severity    string `json:"severity,omitempty"` // low | medium | high (optional for now)
	Source      string `json:"source,omitempty"`   // ai | community | system
}
