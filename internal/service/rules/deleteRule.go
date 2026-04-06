package rules

import (
	"major/pkg/customErrors/serviceErrors"
)

func (s *RulesService) DeleteRule(idRule string) error {
	s.checker.DeleteExpression(idRule)
	err := s.repo.RulesRepository.DeleteRule(idRule)
	if err != nil {
		return serviceErrors.MapError(err)
	}
	return nil
}
