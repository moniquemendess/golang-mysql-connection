package services

import (
	"database/sql"
	"io"
	"log"
	"testing"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

type MockResult struct{}

func (r MockResult) LastInsertId() (int64, error) { return 0, nil }
func (r MockResult) RowsAffected() (int64, error) { return 0, nil }

func TestCreateAuthor(t *testing.T) {
	mockDB := &models.MockDB{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			// Verifica se a query e o argumento estão corretos
			if query != "INSERT INTO authors (name) VALUES (?)" {
				t.Errorf("Query esperada 'INSERT INTO authors (name) VALUES (?)', mas recebeu '%s'", query)
			}
			if args[0] != "Monique" {
				t.Errorf("Esperado nome 'Monique', mas recebeu '%v'", args[0])
			}
			return MockResult{}, nil
		},
	}

	// Redireciona o log para evitar saída durante o teste
	log.SetOutput(io.Discard)

	// Executa o teste
	CreateAuthor(mockDB, "Monique")
}
