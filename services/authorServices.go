package services

import (
	"log"
	"strconv"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

func CreateAuthor(db models.DatabaseExecutor, name string) error {
	newAuthor := models.Author{Name: name}

	if err := models.InserirAutor(db, newAuthor); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Author inserted successfully!")

	return nil
}

func DeleteAuthorID(db models.DatabaseExecutor, id string) error {
	// converter string para int
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err)
		return err
	}

	if err := models.DeleteAuthor(db, intID); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Author deleted successfully!")
	return nil
}
