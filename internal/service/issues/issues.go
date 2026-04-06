package issues

import (
	"major/internal/models"
	"major/internal/repository"
)

type IssuesService struct {
	conf *models.Config
	repo *repository.Repository
}

func NewIssuesService(conf *models.Config, repo *repository.Repository) *IssuesService {
	return &IssuesService{conf: conf, repo: repo}
}
