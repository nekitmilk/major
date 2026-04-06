package rules

import (
	"context"
	"database/sql"
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *RuleRepository) GetRuleFullDataByID(idRule string) (*models.RuleFullData, error) {
	query := fmt.Sprintf(`
		SELECT 
			r.id,
			r.name,
			r.description,
			r.expression,
			r.enabled,
			i.id AS issue_id,
			i.name AS issue_name,
			i.description AS issue_description,
			i.recommendation AS issue_recommendation,
			s.id AS severity_id,
			s.name AS severity_name,
			s.level AS severity_level,
			s.description AS severity_description
		FROM %s r
		JOIN %s i ON r.issue_id = i.id
		JOIN %s s ON i.severity_id = s.id
		WHERE r.id = $1`,
		constRepository.RULES_TABLE_NAME,
		constRepository.ISSUES_TABLE_NAME,
		constRepository.SEVERITIES_TABLE_NAME,
	)

	row := r.DB.QueryRowContext(context.Background(), query, idRule)

	var rf models.RuleFullData
	var issueFull models.IssueFullData
	var severity models.Severity

	err := row.Scan(
		&rf.Id,
		&rf.Name,
		&rf.Description,
		&rf.Expression,
		&rf.Enabled,
		&issueFull.Id,
		&issueFull.Name,
		&issueFull.Description,
		&issueFull.Recommendation,
		&severity.Id,
		&severity.Name,
		&severity.Level,
		&severity.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, externalServiceErrorsModels.NewNotFoundError(fmt.Errorf("rule not found"), idRule)
		}
		return nil, postgres.MapError(err)
	}

	issueFull.Severity = severity
	rf.Issue = issueFull

	return &rf, nil
}
