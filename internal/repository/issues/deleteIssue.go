package issues

import (
	"fmt"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (ir *IssuesRepository) DeleteIssue(idIssue string) error {
	dataDelete, err := ir.DB.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		constRepository.ISSUES_TABLE_NAME),
		idIssue,
	)

	if err != nil {
		return postgres.MapError(err)
	}

	rowsAffected, _ := dataDelete.RowsAffected()

	if rowsAffected == 0 {
		return externalServiceErrorsModels.NewNotFoundError(fmt.Errorf("issue not found"), "")
	}

	return nil
}
