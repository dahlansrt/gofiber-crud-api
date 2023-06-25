package book

import (
	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
)

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	RemoveBook(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}

func (s *service) FetchBooks() (*[]presenter.Book, error) {
	return s.repository.ReadBook()
}

func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdateBook(book)
}

func (s *service) RemoveBook(ID string) error {
	return s.repository.DeleteBook(ID)
}
