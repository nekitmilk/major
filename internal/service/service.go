package service

import (
	"major/internal/models"
	"major/internal/repository"
	"major/internal/service/check"
	"major/internal/service/issues"
	"major/internal/service/rules"
	"major/internal/service/severities"
	"major/pkg/celchecker"
)

type SeverityService interface {
	CreateSeverity(severity *models.CreateSeverityRequest) (string, error)
	DeleteSeverity(idSeverity string) error
	GetAllSeverities() ([]models.Severity, error)
}

type IssueService interface {
	CreateIssue(issue *models.CreateIssueRequest) (string, error)
	DeleteIssue(id string) error
	GetAllIssues() ([]models.Issue, error)
}

type RulesService interface {
	CreateRule(dataRule *models.CreateRuleRequest) (string, error)
	DeleteRule(idRule string) error
	GetAllRules() ([]models.Rule, error)
	GetAllRulesFullData() ([]models.RuleFullData, error)
	GetRuleFullDataByID(idRule string) (*models.RuleFullData, error)
}

type CheckService interface {
	CheckConfig(config map[string]interface{}) ([]string, []error)
	GetConfigSummary(configReq *models.CheckRequest) (*models.CheckResponseFullData, error)
}

type Service struct {
	SeverityService
	IssueService
	RulesService
	CheckService
}

func NewService(conf *models.Config, repo *repository.Repository) *Service {
	checker := celchecker.NewChecker()

	allRules, _ := repo.RulesRepository.GetAllRules()
	expressions := make(map[string]string, len(allRules))
	for _, rule := range allRules {
		expressions[rule.Id] = rule.Expression
	}
	checker.LoadExpressions(expressions)

	return &Service{
		SeverityService: severities.NewSeverityService(repo),
		IssueService:    issues.NewIssuesService(conf, repo),
		RulesService:    rules.NewRulesService(repo, checker),
		CheckService:    check.NewCheckService(checker, repo),
	}
}
