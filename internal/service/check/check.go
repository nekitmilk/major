package check

import (
	"major/internal/repository"
	"major/pkg/celchecker"
)

type CheckService struct {
	checker *celchecker.Checker
	repo    *repository.Repository
}

func NewCheckService(checker *celchecker.Checker, repo *repository.Repository) *CheckService {
	return &CheckService{
		checker: checker,
		repo:    repo,
	}
}
