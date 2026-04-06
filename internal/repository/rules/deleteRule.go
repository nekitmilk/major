package rules

import (
	"fmt"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *RuleRepository) DeleteRule(idRule string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constRepository.RULES_TABLE_NAME)
	result, err := r.DB.Exec(query, idRule)
	if err != nil {
		return postgres.MapError(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return externalServiceErrorsModels.NewNotFoundError(fmt.Errorf("rule not found"), idRule)
	}
	return nil
}
