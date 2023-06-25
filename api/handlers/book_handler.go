package handlers

import (
	"errors"
	"net/http"

	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/book"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AddBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(errors.New("please specify the Title")))
		}

		result, err := service.InsertBook(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(presenter.BookSuccessResponse(result))
	}
}

func UpdateBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(errors.New("please specify the Title")))
		}

		result, err := service.UpdateBook(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(presenter.BookSuccessResponse(result))
	}
}

func RemoveBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		bookID := requestBody.ID
		err = service.RemoveBook(bookID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetBooks(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchBooks()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(presenter.BooksSuccessResponse(fetched))
	}
}
