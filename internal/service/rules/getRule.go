package rules

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (s *RulesService) GetRuleFullDataByID(idRule string) (*models.RuleFullData, error) {
	ruleFull, err := s.repo.RulesRepository.GetRuleFullDataByID(idRule)
	if err != nil {
		return nil, serviceErrors.MapError(err)
	}
	return ruleFull, nil
}
