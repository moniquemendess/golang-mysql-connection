package controllers

import (
	"database/sql"
	"log"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

func CreateAuthor(db *sql.DB, name string) {
	newAuthor := models.Author{Name: name}

	if err := models.InserirAutor(db, newAuthor); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Author inserted successfully!")
	}
}
