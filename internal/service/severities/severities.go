package severities

import "major/internal/repository"

type SeverityService struct {
	repo *repository.Repository
}

func NewSeverityService(repo *repository.Repository) *SeverityService {
	return &SeverityService{repo: repo}
}
