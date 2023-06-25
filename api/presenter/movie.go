package presenter

import (
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title      string             `json:"title" bson:"title"`
	Director   string             `json:"director" bson:"director"`
	Screenplay string             `json:"screenplay" bson:"screenplay"`
}

func MovieSuccessResponse(data *entities.Movie) *fiber.Map {
	movie := Movie{
		ID:         data.ID,
		Title:      data.Title,
		Director:   data.Director,
		Screenplay: data.Screenplay,
	}
	return &fiber.Map{
		"status": true,
		"data":   movie,
		"error":  nil,
	}
}

func MoviesSuccessResponse(data *[]Movie) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func MovieErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
