package severities

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (s *SeverityService) CreateSeverity(severity *models.CreateSeverityRequest) (string, error) {
	idSeverity, err := s.repo.SeverityRepository.CreateSeverity(severity)
	if err != nil {
		return "", serviceErrors.MapError(err)
	}

	return idSeverity, nil
}
