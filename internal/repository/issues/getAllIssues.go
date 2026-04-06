package issues

import (
	"context"
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (ir *IssuesRepository) GetAllIssues() ([]models.Issue, error) {
	query := fmt.Sprintf(
		`SELECT 
			id, 
			name, 
			description,
			recommendation,
			severity_id
		FROM %s`,
		constRepository.ISSUES_TABLE_NAME,
	)

	ctx := context.Background()
	rows, err := ir.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.MapError(err)
	}

	defer rows.Close()

	var issues []models.Issue

	for rows.Next() {
		var issue models.Issue

		errScan := rows.Scan(
			&issue.Id,
			&issue.Name,
			&issue.Description,
			&issue.Recommendation,
			&issue.SeverityId,
		)

		if errScan != nil {
			return nil, errScan
		}

		issues = append(issues, issue)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return issues, nil
}
