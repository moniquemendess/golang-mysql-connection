package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hola CHRIS")
	})
}
