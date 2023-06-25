package presenter

import (
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Author string             `json:"author" bson:"author"`
}

func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book := Book{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

func BooksSuccessResponse(data *[]Book) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
