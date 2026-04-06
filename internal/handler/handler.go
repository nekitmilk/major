package handler

import (
	"major/internal/handler/check"
	"major/internal/handler/health"
	"major/internal/handler/issues"
	"major/internal/handler/rules"
	"major/internal/handler/severities"
	"major/internal/models"
	"major/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthHandler interface {
	HealthCheck(ctx *gin.Context)
}

type SeverityHandler interface {
	CreateSeverity(ctx *gin.Context)
	DeleteSeverity(ctx *gin.Context)
	GetAllSeverities(ctx *gin.Context)
}

type IssueHandler interface {
	CreateIssue(ctx *gin.Context)
	DeleteIssue(ctx *gin.Context)
	GetAllIssues(ctx *gin.Context)
}

type RulesHandler interface {
	CreateRule(ctx *gin.Context)
	DeleteRule(ctx *gin.Context)
	GetAllRules(ctx *gin.Context)
	GetAllRulesFullData(ctx *gin.Context)
	GetRuleFullDataByID(ctx *gin.Context)
}

type CheckHandler interface {
	//CheckConfig(ctx *gin.Context)
	GetConfigSummary(ctx *gin.Context)
}

type Handler struct {
	log      *logrus.Logger
	conf     *models.Config
	services *service.Service
	HealthHandler
	SeverityHandler
	IssueHandler
	RulesHandler
	CheckHandler
}

func NewHandler(log *logrus.Logger, conf *models.Config, services *service.Service) *Handler {
	return &Handler{
		log:             log,
		conf:            conf,
		HealthHandler:   health.NewHealthHandler(log),
		SeverityHandler: severities.NewSeveritiesHandler(log, conf, services),
		IssueHandler:    issues.NewIssuesHandler(log, conf, services),
		RulesHandler:    rules.NewRulesHandler(log, conf, services),
		CheckHandler:    check.NewCheckHandler(log, conf, services),
	}
}
