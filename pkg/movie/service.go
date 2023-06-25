package movie

import (
	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
)

type Service interface {
	InsertMovie(movie *entities.Movie) (*entities.Movie, error)
	FetchMovies() (*[]presenter.Movie, error)
	UpdateMovie(movie *entities.Movie) (*entities.Movie, error)
	RemoveMovie(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertMovie(movie *entities.Movie) (*entities.Movie, error) {
	return s.repository.CreateMovie(movie)
}

func (s *service) FetchMovies() (*[]presenter.Movie, error) {
	return s.repository.ReadMovie()
}

func (s *service) UpdateMovie(movie *entities.Movie) (*entities.Movie, error) {
	return s.repository.UpdateMovie(movie)
}

func (s *service) RemoveMovie(ID string) error {
	return s.repository.DeleteMovie(ID)
}
