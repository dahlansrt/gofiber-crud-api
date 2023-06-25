package routes

import (
	"github.com/dahlansrt/gofiber-crud-api/api/handlers"
	"github.com/dahlansrt/gofiber-crud-api/pkg/book"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/book", handlers.GetBooks(service))
	app.Post("/book", handlers.AddBook(service))
	app.Put("/book", handlers.UpdateBook(service))
	app.Delete("/book", handlers.RemoveBook(service))
}
