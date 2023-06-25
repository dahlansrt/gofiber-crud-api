package employee

import (
	"context"
	"time"

	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateEmployee(employee *entities.Employee) (*entities.Employee, error)
	ReadEmployee() (*[]presenter.Employee, error)
	UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	DeleteEmployee(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	_, err := r.Collection.InsertOne(context.Background(), employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *repository) ReadEmployee() (*[]presenter.Employee, error) {
	var employees []presenter.Employee

	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var employee presenter.Employee
		_ = cursor.Decode(&employee)
		employees = append(employees, employee)
	}

	return &employees, nil
}

func (r *repository) UpdateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	employee.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": employee.ID}, bson.M{"$set": employee})
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *repository) DeleteEmployee(ID string) error {
	employeeID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": employeeID})
	if err != nil {
		return err
	}

	return nil
}
