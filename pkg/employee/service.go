package employee

import (
	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
)

type Service interface {
	InsertEmployee(employee *entities.Employee) (*entities.Employee, error)
	FetchEmployees() (*[]presenter.Employee, error)
	UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	RemoveEmployee(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertEmployee(employee *entities.Employee) (*entities.Employee, error) {
	return s.repository.CreateEmployee(employee)
}

func (s *service) FetchEmployees() (*[]presenter.Employee, error) {
	return s.repository.ReadEmployee()
}

func (s *service) UpdateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	return s.repository.UpdateEmployee(employee)
}

func (s *service) RemoveEmployee(ID string) error {
	return s.repository.DeleteEmployee(ID)
}
