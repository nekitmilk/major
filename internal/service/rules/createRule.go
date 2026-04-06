package rules

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
	"major/pkg/customErrors/serviceErrors/serviceErrorsModels"
)

func (s *RulesService) CreateRule(dataRule *models.CreateRuleRequest) (string, error) {
	id, err := s.repo.RulesRepository.CreateRule(dataRule)
	if err != nil {
		return "", serviceErrors.MapError(err)
	}
	errAddExpression := s.checker.AddExpression(id, dataRule.Expression)
	if errAddExpression != nil {
		_ = s.repo.RulesRepository.DeleteRule(id) // TODO: реализовать поддержку транзакций постгри
		return "", serviceErrorsModels.NewCelExpressionError(errAddExpression, "")
	}
	return id, nil
}
