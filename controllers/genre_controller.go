package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/moniquemendess/golang-mysql-connection/services"
)

type GenreController struct {
	Service services.GenresServices // GenresServices é a interface
}

func NewGenreController(service services.GenresServices) *GenreController {
	return &GenreController{Service: service}
}

func (controller *GenreController) CreateGenre(c fiber.Ctx) error {
	name := c.Params("name")
	genre, err := controller.Service.CreateGenre(name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(genre)
}
