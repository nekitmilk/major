package rules

import (
	"context"
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *RuleRepository) GetAllRules() ([]models.Rule, error) {
	query := fmt.Sprintf(`
		SELECT id, name, description, issue_id, expression, enabled
		FROM %s`, constRepository.RULES_TABLE_NAME,
	)

	rows, err := r.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, postgres.MapError(err)
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		if err := rows.Scan(&rule.Id, &rule.Name, &rule.Description, &rule.IssueId, &rule.Expression, &rule.Enabled); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rules, nil
}

func (r *RuleRepository) GetAllRulesFullData() ([]models.RuleFullData, error) {
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
		JOIN %s s ON i.severity_id = s.id`,
		constRepository.RULES_TABLE_NAME,
		constRepository.ISSUES_TABLE_NAME,
		constRepository.SEVERITIES_TABLE_NAME,
	)

	rows, err := r.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, postgres.MapError(err)
	}
	defer rows.Close()

	var rulesFull []models.RuleFullData
	for rows.Next() {
		var rf models.RuleFullData
		var issueFull models.IssueFullData
		var severity models.Severity

		errScan := rows.Scan(
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
		if errScan != nil {
			return nil, errScan
		}

		issueFull.Severity = severity
		rf.Issue = issueFull
		rulesFull = append(rulesFull, rf)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rulesFull, nil
}
