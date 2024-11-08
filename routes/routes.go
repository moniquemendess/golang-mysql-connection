package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/moniquemendess/golang-mysql-connection/models"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

func Setup(app *fiber.App, db *sql.DB) {
	app.Post("/create-author/:name", func(c fiber.Ctx) error {
		return services.CreateAuthor(db, c.Params("name"))
	})
	app.Post("/create-genre/:name", func(c fiber.Ctx) error {
		return services.CreateGenre(db, c.Params("name"))

	})

	app.Post("/update-author/:id/:name", func(c fiber.Ctx) error {
		id := c.Params("id")
		name := c.Params("name")
		return services.UpdateAuthorID(db, id, name)
	})

	app.Post("/update-genre/:id/:name", func(c fiber.Ctx) error {
		id := c.Params("id")
		name := c.Params("name")
		return services.UpdateGenreID(db, id, name)
	})

	app.Delete("/delete-author/:id", func(c fiber.Ctx) error {
		return services.DeleteAuthorID(db, c.Params("id"))
	})

	app.Delete("/delete-genre/:id", func(c fiber.Ctx) error {
		return services.DeleteGenreID(db, c.Params("id"))
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

	app.Get("/genres", func(c fiber.Ctx) error {
		genres, err := models.FetchAllGenres(db)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch genres"})
		}
		return c.JSON(genres)
	})

}
