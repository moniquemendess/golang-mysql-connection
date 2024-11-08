package services

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

func CreateGenre(db *sql.DB, name string) error {
	newGenre := models.Genres{Name: name}

	if err := models.InsertGenre(db, newGenre); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Genre inserted successfully!")

	return nil
}

func UpdateGenreID(db *sql.DB, id, name string) error {

	// converter string para int
	intID, err := strconv.Atoi(id)
	if err != nil {

		return err
	}
	// llamando la funcion UpdateGenre y pasando el id y el nombre
	if err := models.UpdateGenre(db, intID, name); err != nil {

		return err
	}
	log.Println("Genre modified successfully!")
	return nil
}

func GetGenres(db *sql.DB) ([]models.Genres, error) {
	return models.FetchAllGenres(db)
}

func DeleteGenreID(db *sql.DB, id string) error {
	// converter string para int
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err)
		return err
	}

	if err := models.DeleteGenre(db, intID); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Genre deleted successfully!")
	return nil
}
