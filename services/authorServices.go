package services

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

func CreateAuthor(db *sql.DB, name string) error {
	newAuthor := models.Author{Name: name}

	if err := models.InserirAutor(db, newAuthor); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Author inserted successfully!")

	return nil
}

func UpdateAuthorID(db *sql.DB, id, name string) error {

	// converter string para int
	intID, err := strconv.Atoi(id)
	if err != nil {

		return err
	}
	// llamando la funcion UpdateAuthor y pasando el id y el nombre
	if err := models.UpdateAuthor(db, intID, name); err != nil {

		return err
	}
	log.Println("Author modified successfully!")
	return nil
}

// limit y offset(indice inicial para la consulta)
func QueryAuthorsWithPagination(db *sql.DB, limit, offset int) ([]models.Author, error) {
	query := "SELECT id, name FROM authors LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author // Inicializa la fatia de autores como vacia
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author) // Adiciona o autor à fatia
	}

	return authors, nil
}

func DeleteAuthorID(db *sql.DB, id string) error {
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
