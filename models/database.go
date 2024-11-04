package models

import "database/sql"

// Definicíon de la interface
type DatabaseExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}
