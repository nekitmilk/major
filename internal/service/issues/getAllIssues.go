package issues

import (
	"major/internal/models"
	"major/pkg/customErrors/serviceErrors"
)

func (is *IssuesService) GetAllIssues() ([]models.Issue, error) {
	issues, err := is.repo.IssueRepository.GetAllIssues()
	if err != nil {
		return nil, serviceErrors.MapError(err)
	}
	return issues, nil
}
