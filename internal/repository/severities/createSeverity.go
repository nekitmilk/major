package severities

import (
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *SeverityRepository) CreateSeverity(severity *models.CreateSeverityRequest) (string, error) {
	var id string
	query := fmt.Sprintf(`
		INSERT INTO %s (name, level, description)
		VALUES ($1, $2, $3)
		RETURNING id`,
		constRepository.SEVERITIES_TABLE_NAME,
	)

	err := r.DB.QueryRow(query, severity.Name, severity.Level, severity.Description).Scan(&id)
	if err != nil {
		return "", postgres.MapError(err)
	}
	return id, nil
}
