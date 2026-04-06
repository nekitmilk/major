package rules

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (s *RulesService) GetAllRules() ([]models.Rule, error) {
	rules, err := s.repo.RulesRepository.GetAllRules()
	if err != nil {
		return nil, serviceErrors.MapError(err)
	}
	return rules, nil
}

func (s *RulesService) GetAllRulesFullData() ([]models.RuleFullData, error) {
	rulesFull, err := s.repo.RulesRepository.GetAllRulesFullData()
	if err != nil {
		return nil, serviceErrors.MapError(err)
	}
	return rulesFull, nil
}
