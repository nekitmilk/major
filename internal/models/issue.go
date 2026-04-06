package models

type Issue struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
	SeverityId     string `json:"severity_id"`
}

type CreateIssueRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
	SeverityId     string `json:"severity_id"`
}

type IssueFullData struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Recommendation string   `json:"recommendation"`
	Severity       Severity `json:"severities"`
}

type IssueResponse struct {
	Id string `json:"id"`
}
