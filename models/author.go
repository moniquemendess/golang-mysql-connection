package models

// struct author
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// insert author function
func InserirAutor(db DatabaseExecutor, author Author) error {
	query := "INSERT INTO authors (name) VALUES (?)"
	_, err := db.Exec(query, author.Name)
	return err
}

// delete user
func DeleteUser(db DatabaseExecutor, id int) error {
	query := "DELETE FROM users Where id = ?"
	_, err := db.Exec(query, id)
	return err
}
