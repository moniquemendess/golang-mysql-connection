package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
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
	// routes.SetupRoutes(app, db)

	// Start the server on port 3000
	log.Println("Server is running onhttp://localhost:3001")
	log.Fatal(app.Listen(":3001"))
}
