package severities

import "major/pkg/customErrors/serviceErrors"

func (s *SeverityService) DeleteSeverity(idSeverity string) error {
	err := s.repo.SeverityRepository.DeleteSeverity(idSeverity)
	if err != nil {
		return serviceErrors.MapError(err)
	}
	return nil
}
