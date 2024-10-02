package mysql

import "database/sql"

type sqlDB struct {
	*sql.DB
}

func NewSqlDB(driver, connStr string) (repo *sqlDB, err error) {
	db, err := sql.Open(driver, connStr)

	return &sqlDB{
		db,
	}, err
}
