package models

type CheckRequest struct {
	Config string `json:"config"`
}

type CheckResponse struct {
	RuleIds []string `json:"rule_ids"`
	Errors  []string `json:"errors"`
}

type CheckResponseFullData struct {
	Rules  []*RuleFullData `json:"rules"`
	Errors []string        `json:"errors"`
}
