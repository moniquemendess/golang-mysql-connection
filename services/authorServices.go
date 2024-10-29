package services

import (
	"log"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

func CreateAuthor(db models.DatabaseExecutor, name string) {
	newAuthor := models.Author{Name: name}

	if err := models.InserirAutor(db, newAuthor); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Author inserted successfully!")
	}
}
