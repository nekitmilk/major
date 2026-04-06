package repository

import (
	"database/sql"
	"major/internal/models"
	issue "major/internal/repository/issues"
	rule "major/internal/repository/rules"
	"major/internal/repository/severities"
)

type SeverityRepository interface {
	CreateSeverity(severity *models.CreateSeverityRequest) (string, error)
	DeleteSeverity(idSeverity string) error
	GetAllSeverities() ([]models.Severity, error)
}

type IssueRepository interface {
	CreateIssue(issue *models.CreateIssueRequest) (string, error)
	DeleteIssue(id string) error
	GetAllIssues() ([]models.Issue, error)
}

type RulesRepository interface {
	CreateRule(dataRule *models.CreateRuleRequest) (string, error)
	DeleteRule(idRule string) error
	GetAllRules() ([]models.Rule, error)
	GetAllRulesFullData() ([]models.RuleFullData, error)
	GetRuleFullDataByID(idRule string) (*models.RuleFullData, error)
}

type Repository struct {
	db *sql.DB
	SeverityRepository
	IssueRepository
	RulesRepository
}

func NewRepository(clientSQL *sql.DB) *Repository {
	return &Repository{
		SeverityRepository: severities.NewSeverityRepository(clientSQL),
		IssueRepository:    issue.NewIssuesRepository(clientSQL),
		RulesRepository:    rule.NewRuleRepository(clientSQL),
	}
}
