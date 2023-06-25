package routes

import (
	"github.com/dahlansrt/gofiber-crud-api/api/handlers"
	"github.com/dahlansrt/gofiber-crud-api/pkg/employee"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRouter(app fiber.Router, service employee.Service) {
	app.Get("/employees", handlers.GetEmployees(service))
	app.Post("/employees", handlers.AddEmployee(service))
	app.Put("/employees", handlers.UpdateEmployee(service))
	app.Delete("/employees", handlers.RemoveEmployee(service))
}