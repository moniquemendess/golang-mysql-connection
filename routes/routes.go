package routes

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

func Setup(app *fiber.App, db *sql.DB) {
	app.Post("/create-author/:name", func(c fiber.Ctx) error {
		return services.CreateAuthor(db, c.Params("name"))
	})

	app.Post("/update-author/:id/:name", func(c fiber.Ctx) error {
		id := c.Params("id")
		name := c.Params("name")
		return services.UpdateAuthorID(db, id, name)
	})

	app.Delete("/delete-author/:id", func(c fiber.Ctx) error {
		return services.DeleteAuthorID(db, c.Params("id"))
	})

	app.Get("/authors/:limit/:offset", func(c fiber.Ctx) error {
		limit, err := strconv.Atoi(c.Params("limit"))
		if err != nil {
			return c.Status(400).SendString("Invalid limit parameter")
		}
		offset, err := strconv.Atoi(c.Params("offset"))
		if err != nil {
			return c.Status(400).SendString("Invalid offset parameter")
		}

		authors, err := services.QueryAuthorsWithPagination(db, limit, offset)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(authors)
	})

}
