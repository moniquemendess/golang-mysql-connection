package repository

import (
	"database/sql"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

// Struct de conexao com a base de dados
type GenresRepository struct {
	DB *sql.DB
}

// Funcao construtura cria e retorna uma nova instancia do repositorio (GenresRepository)
func NewGenresRepository(db *sql.DB) models.Repository[models.Genres] {
	return &GenresRepository{DB: db}
}

func (repo *GenresRepository) Create(genres models.Genres) (models.Genres, error) {

	query := "INSERT INTO genres (name) VALUES (?)"
	result, err := repo.DB.Exec(query, genres.Name) // executa a query e substitui ? por genres.Names
	if err != nil {                                 // verifica se houve erro, se sim retorna um valor vazio em models.Genres

		return models.Genres{}, err
	}
	id, err := result.LastInsertId() //Obtém o ID do último registro inserido no banco de dados
	if err != nil {                  //Se houver um erro, retorna um valor vazio de models.Genres
		return models.Genres{}, err
	}
	genres.ID = int(id) // Atualiza o campo ID do objeto genres con o ID receém obtido do banco de dados.
	return genres, nil
}

func (repo *GenresRepository) Update(genres models.Genres) error {
	query := "Update genres SET name = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, genres.Name, genres.ID)
	return err
}

func (repo *GenresRepository) Delete(id int) error {
	query := "DELETE FROM genres Where id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}
