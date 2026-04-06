package severities

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (s *SeverityService) GetAllSeverities() ([]models.Severity, error) {
	severities, err := s.repo.SeverityRepository.GetAllSeverities()
	if err != nil {
		return nil, serviceErrors.MapError(err)
	}
	return severities, nil
}
