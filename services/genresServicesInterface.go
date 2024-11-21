package services

import (
	"errors"

	"github.com/moniquemendess/golang-mysql-connection/models"
)

type GenresServices interface {
	CreateGenre(name string) (models.Genres, error)
	UpdateGenre(id int, name string) error
	DeleteGenre(id int) error
}

// Struct de Interface Genérica da interface GenresServices
type GenresServicesImpl struct {
	Repo models.Repository[models.Genres]
}

func NewGenresServices(repo models.Repository[models.Genres]) GenresServices {
	return &GenresServicesImpl{Repo: repo}
}

func (s *GenresServicesImpl) CreateGenre(name string) (models.Genres, error) {
	if len(name) == 0 { // verifica se name é uma string vazia e retorna um erro
		return models.Genres{}, errors.New("genre name cannout be empty")
	}
	genre := models.Genres{Name: name} // criando uma instancia da struct de models.Genres com o campo Name
	return s.Repo.Create(genre)        //passando a instancia genres para salvar no BD
}

func (s *GenresServicesImpl) UpdateGenre(id int, name string) error {
	genre := models.Genres{ID: id, Name: name} // criando uma instancia da struct de models.Genres com o campo id/name passados por parametros para atualizar.
	return s.Repo.Update(genre)                // passando a instancia genres para salvar no BD
}

func (s *GenresServicesImpl) DeleteGenre(id int) error {
	return s.Repo.Delete(id) // id por params
}
