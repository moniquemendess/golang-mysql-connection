package services

import (
	"log"

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
