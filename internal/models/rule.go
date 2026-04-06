package models

type Rule struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IssueId     string `json:"issue_id"`
	Expression  string `json:"expression"`
	Enabled     bool   `json:"enabled"`
}

type CreateRuleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IssueId     string `json:"issue_id"`
	Expression  string `json:"expression"`
	Enabled     bool   `json:"enabled"`
}

type RuleFullData struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Issue       IssueFullData `json:"issue"`
	Expression  string        `json:"expression"`
	Enabled     bool          `json:"enabled"`
}

type RuleResponse struct {
	Id string `json:"id"`
}
