package routes

import (
	"database/sql"
	"fmt"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

func SetupRoutes(app fiber.App, db *sql.DB) {
	app.Post("/create-author", func(c fiber.Ctx) error {
		fmt.Println("hit")
		services.CreateAuthor(db, "Test") // this function should return at least an error
		return c.SendStatus(200)
	})
}
