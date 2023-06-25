package handlers

import (
	"errors"
	"net/http"

	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/employee"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AddEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		if requestBody.Name == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New("please specify the Name")))
		}

		result, err := service.InsertEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

func UpdateEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		if requestBody.Name == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New("please specify the Name")))
		}

		result, err := service.UpdateEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

func RemoveEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		employeeID := requestBody.ID
		err = service.RemoveEmployee(employeeID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetEmployees(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchEmployees()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}

		return c.JSON(presenter.EmployeesSuccessResponse(fetched))
	}
}
