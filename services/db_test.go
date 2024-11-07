package services

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql" // ou outro driver que você estiver usando
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// Criar um novo pool de conexões do Docker
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Criar um container para o banco de dados
	resource, err := pool.Run("mysql", "latest", []string{
		"MYSQL_ROOT_PASSWORD=root",
		"MYSQL_DATABASE=testdb",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Expondo a porta
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:"+resource.GetPort("3306/tcp")+")/testdb")
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Executar os testes
	m.Run()

	// Fechar o container após os testes
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
