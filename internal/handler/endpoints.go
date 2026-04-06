package handler

import (
	"major/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (handler *Handler) InitRoutes(conf *models.Config) *gin.Engine {
	routes := gin.New()

	routes.Use(handler.CorsMiddleware())

	routes.StaticFS("/docs", http.Dir("./docs"))

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/docs/swagger.json")))

	routes.GET("/health", handler.HealthHandler.HealthCheck)

	api := routes.Group("/api/v1")
	{
		check := api.Group("/check")
		{
			check.POST("", handler.CheckHandler.GetConfigSummary)
		}

		// Severities
		severities := api.Group("/severities")
		{
			severities.POST("", handler.SeverityHandler.CreateSeverity)
			severities.DELETE("/:severityId", handler.SeverityHandler.DeleteSeverity)
			severities.GET("", handler.SeverityHandler.GetAllSeverities)
		}

		// Issues
		issues := api.Group("/issues")
		{
			issues.POST("", handler.IssueHandler.CreateIssue)
			issues.DELETE("/:issueId", handler.IssueHandler.DeleteIssue)
			issues.GET("", handler.IssueHandler.GetAllIssues)
		}

		// Rules
		rules := api.Group("/rules")
		{
			rules.POST("", handler.RulesHandler.CreateRule)
			rules.DELETE("/:ruleId", handler.RulesHandler.DeleteRule)
			rules.GET("", handler.RulesHandler.GetAllRules)

			// Full data routes
			rules.GET("/full", handler.RulesHandler.GetAllRulesFullData)
			rules.GET("/full/:ruleId", handler.RulesHandler.GetRuleFullDataByID)
		}
	}

	return routes
}
