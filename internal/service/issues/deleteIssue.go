package issues

import "major/pkg/customErrors/serviceErrors"

func (is *IssuesService) DeleteIssue(idIssue string) error {
	err := is.repo.IssueRepository.DeleteIssue(idIssue)
	if err != nil {
		return serviceErrors.MapError(err)
	}

	return nil
}
