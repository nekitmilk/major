package issues

import (
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (ir *IssuesRepository) CreateIssue(dataIssue *models.CreateIssueRequest) (string, error) {
	var id string
	query := fmt.Sprintf(`
		INSERT INTO %s (
			name,
			description,
			recommendation,
			severity_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		constRepository.ISSUES_TABLE_NAME,
	)

	err := ir.DB.QueryRow(query,
		dataIssue.Name,
		dataIssue.Description,
		dataIssue.Recommendation,
		dataIssue.SeverityId,
	).Scan(&id)

	if err != nil {
		return "", postgres.MapError(err)
	}
	return id, nil
}
