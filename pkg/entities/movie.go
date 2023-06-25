package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title      string             `json:"title" bson:"title"`
	Director   string             `json:"director" bson:"director"`
	Screenplay string             `json:"screenplay" bson:"screenplay"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}
