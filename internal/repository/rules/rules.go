package rules

import "database/sql"

type RuleRepository struct {
	DB *sql.DB
}

func NewRuleRepository(db *sql.DB) *RuleRepository {
	return &RuleRepository{
		DB: db,
	}
}
