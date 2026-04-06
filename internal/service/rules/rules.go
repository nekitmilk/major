package rules

import (
	"major/internal/repository"
	"major/pkg/celchecker"
)

type RulesService struct {
	repo    *repository.Repository
	checker *celchecker.Checker
}

func NewRulesService(repo *repository.Repository, checker *celchecker.Checker) *RulesService {
	return &RulesService{
		repo:    repo,
		checker: checker,
	}
}
