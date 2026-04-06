package issues

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (is *IssuesService) CreateIssue(issue *models.CreateIssueRequest) (string, error) {
	idIssue, err := is.repo.IssueRepository.CreateIssue(issue)
	if err != nil {
		return "", serviceErrors.MapError(err)
	}

	return idIssue, nil
}
