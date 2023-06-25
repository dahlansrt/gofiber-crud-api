package handlers

import (
	"errors"
	"net/http"

	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/dahlansrt/gofiber-crud-api/pkg/movie"
	"github.com/gofiber/fiber/v2"
)

func AddMovie(service movie.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Movie
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(errors.New("please specify the Title")))
		}

		result, err := service.InsertMovie(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		return c.JSON(presenter.MovieSuccessResponse(result))
	}
}

func UpdateMovie(service movie.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Movie
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(errors.New("please specify the Title")))
		}

		result, err := service.UpdateMovie(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		return c.JSON(presenter.MovieSuccessResponse(result))
	}
}

func RemoveMovie(service movie.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		movieID := requestBody.ID
		err = service.RemoveMovie(movieID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetMovies(service movie.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchMovies()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}

		return c.JSON(presenter.MoviesSuccessResponse(fetched))
	}
}
