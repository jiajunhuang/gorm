package gorm

import "database/sql"

// SQLCommon is the minimal database connection functionality gorm requires.  Implemented by *sql.DB.
// *sql.DB 实现了这些方法
type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type sqlDb interface {
	Begin() (*sql.Tx, error)
}

type sqlTx interface {
	Commit() error
	Rollback() error
}
