package severities

import (
	"context"
	"fmt"
	"major/internal/models"
	"major/internal/repository/constRepository"
	"major/pkg/customErrors/externalServiceErrors/postgres"
)

func (r *SeverityRepository) GetAllSeverities() ([]models.Severity, error) {
	query := fmt.Sprintf(`
		SELECT id, name, level, description
		FROM %s`, constRepository.SEVERITIES_TABLE_NAME,
	)

	rows, err := r.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, postgres.MapError(err)
	}
	defer rows.Close()

	var severities []models.Severity
	for rows.Next() {
		var s models.Severity
		if errScan := rows.Scan(&s.Id, &s.Name, &s.Level, &s.Description); errScan != nil {
			return nil, errScan
		}
		severities = append(severities, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return severities, nil
}
