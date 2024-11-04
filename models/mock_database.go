// models/mock_database.go
package models

import "database/sql"

type MockDB struct {
	ExecFunc func(query string, args ...interface{}) (sql.Result, error)
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.ExecFunc(query, args...)
}
