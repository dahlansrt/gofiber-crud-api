package presenter

import (
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func EmployeeSuccessResponse(data *entities.Employee) *fiber.Map {
	employee := Employee{
		ID:     data.ID,
		Name:   data.Name,
		Salary: data.Salary,
		Age:    data.Age,
	}
	return &fiber.Map{
		"status": true,
		"data":   employee,
		"error":  nil,
	}
}

func EmployeesSuccessResponse(data *[]Employee) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func EmployeeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
