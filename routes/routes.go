package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/moniquemendess/golang-mysql-connection/controllers"
	"github.com/moniquemendess/golang-mysql-connection/repository"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

func Setup(app *fiber.App, db *sql.DB) {

	genresRepo := repository.NewGenresRepository(db)
	genresService := services.NewGenresServices(genresRepo)
	genresController := controllers.NewGenreController(genresService)

	app.Post("/create-genre/:name", genresController.CreateGenre)
}
