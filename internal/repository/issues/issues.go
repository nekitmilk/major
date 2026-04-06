package issues

import "database/sql"

type IssuesRepository struct {
	DB *sql.DB
}

func NewIssuesRepository(db *sql.DB) *IssuesRepository {
	return &IssuesRepository{
		DB: db,
	}
}
