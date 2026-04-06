package rules

import (
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *RuleRepository) CreateRule(dataRule *models.CreateRuleRequest) (string, error) {
	var id string
	query := fmt.Sprintf(`
		INSERT INTO %s (name, description, issue_id, expression, enabled)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		constRepository.RULES_TABLE_NAME,
	)

	err := r.DB.QueryRow(query,
		dataRule.Name,
		dataRule.Description,
		dataRule.IssueId,
		dataRule.Expression,
		dataRule.Enabled,
	).Scan(&id)

	if err != nil {
		return "", postgres.MapError(err)
	}
	return id, nil
}
