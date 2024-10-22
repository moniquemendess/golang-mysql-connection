package models

import (
	"database/sql"
)

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
