package models

type Severity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type CreateSeverityRequest struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

type SeverityResponse struct {
	Id string `json:"id"`
}
