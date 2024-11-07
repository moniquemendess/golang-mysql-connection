package models

import "database/sql"

// struct author
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// insert author function
func InserirAutor(db *sql.DB, author Author) error {
	query := "INSERT INTO authors (name) VALUES (?)"
	_, err := db.Exec(query, author.Name)
	return err
}

func UpdateAuthor(db *sql.DB, id int, name string) error {
	query := "Update authors SET name = ? WHERE id = ?"
	_, err := db.Exec(query, name, id) // en la orden de la consulta query
	return err
}

// delete author
func DeleteAuthor(db *sql.DB, id int) error {
	query := "DELETE FROM authors Where id = ?"
	_, err := db.Exec(query, id)
	return err
}
