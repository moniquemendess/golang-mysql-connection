package models

import (
	"database/sql"
)

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// insert genres function
func InsertGenre(db *sql.DB, genres Genres) error {
	query := "INSERT INTO genres (name) VALUES (?)"
	_, err := db.Exec(query, genres.Name)
	return err
}

// update genres
func UpdateGenre(db *sql.DB, id int, name string) error {
	query := "Update genres SET name = ? WHERE id = ?"
	_, err := db.Exec(query, name, id) // en la orden de la consulta query
	return err
}

func FetchAllGenres(db *sql.DB) ([]Genres, error) {
	rows, err := db.Query("SELECT * FROM genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []Genres // inicializa vacío
	for rows.Next() {
		var genre Genres
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, err
		}
		genres = append(genres, genre) // adiciona los generos
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return genres, nil // Retorna a fatia de generos
}

// delete genres
func DeleteGenre(db *sql.DB, id int) error {
	query := "DELETE FROM genres Where id = ?"
	_, err := db.Exec(query, id)
	return err
}
