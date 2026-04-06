package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (sqlDB *SQLDB) InitDB() (*sql.DB, error) {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		sqlDB.conf.Db.Host,
		sqlDB.conf.Db.Port,
		sqlDB.conf.Db.User,
		sqlDB.conf.Db.Password,
		sqlDB.conf.Db.DbName,
		sqlDB.conf.Db.Sslmode,
	)

	db, err := sql.Open(sqlDB.conf.Db.DriverName, connect)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
