package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	fiber "github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Recupera a senha do banco de dados da variável de ambiente
	dbPassword := os.Getenv("DB_PASSWORD")

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:"+dbPassword+"@tcp(127.0.0.1:3306)/librosdb")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Testa a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Configura as rotas
	app.Post("/create-author/:name", func(c fiber.Ctx) error {
		return services.CreateAuthor(db, c.Params("name"))
	})

	app.Delete("/delete-author/:id", func(c fiber.Ctx) error {
		return services.DeleteAuthorID(db, c.Params("id"))
	})

	// Start the server on port 3000
	log.Println("Server is running on http://localhost:3001")
	log.Fatal(app.Listen("localhost:3001"))
}
