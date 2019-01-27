package dbx

import "database/sql"

// Scannable interface defines a set of methods shared by sql.Row and sql.Rows
type Scannable interface {
	Scan(dest ...interface{}) error
}

// Queryable interface defines a set of methods shared by sql.DB and sql.Tx
type Queryable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
