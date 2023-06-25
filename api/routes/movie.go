package routes

import (
	"github.com/dahlansrt/gofiber-crud-api/api/handlers"
	"github.com/dahlansrt/gofiber-crud-api/pkg/movie"
	"github.com/gofiber/fiber/v2"
)

func MovieRouter(app fiber.Router, service movie.Service) {
	app.Get("/movie", handlers.GetMovies(service))
	app.Post("/movie", handlers.AddMovie(service))
	app.Put("/movie", handlers.UpdateMovie(service))
	app.Delete("/movie", handlers.RemoveMovie(service))
}
