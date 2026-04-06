package severities

import (
	"fmt"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *SeverityRepository) DeleteSeverity(idSeverity string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", constRepository.SEVERITIES_TABLE_NAME)
	result, err := r.DB.Exec(query, idSeverity)
	if err != nil {
		return postgres.MapError(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return externalServiceErrorsModels.NewNotFoundError(fmt.Errorf("severities not found"), idSeverity)
	}
	return nil
}
