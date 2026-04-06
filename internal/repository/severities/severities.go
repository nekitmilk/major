package severities

import "database/sql"

type SeverityRepository struct {
	DB *sql.DB
}

func NewSeverityRepository(db *sql.DB) *SeverityRepository {
	return &SeverityRepository{
		DB: db,
	}
}
