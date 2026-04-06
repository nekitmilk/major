package sql

import "major/internal/models"

type SQLDB struct {
	conf *models.Config
}

func NewSQL(conf *models.Config) *SQLDB {
	return &SQLDB{
		conf: conf,
	}
}
