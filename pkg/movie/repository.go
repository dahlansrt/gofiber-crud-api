package movie

import (
	"context"
	"time"

	"github.com/dahlansrt/gofiber-crud-api/api/presenter"
	"github.com/dahlansrt/gofiber-crud-api/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateMovie(movie *entities.Movie) (*entities.Movie, error)
	ReadMovie() (*[]presenter.Movie, error)
	UpdateMovie(movie *entities.Movie) (*entities.Movie, error)
	DeleteMovie(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateMovie(movie *entities.Movie) (*entities.Movie, error) {
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	_, err := r.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *repository) ReadMovie() (*[]presenter.Movie, error) {
	var movies []presenter.Movie

	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var movie presenter.Movie
		_ = cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	return &movies, nil
}

func (r *repository) UpdateMovie(movie *entities.Movie) (*entities.Movie, error) {
	movie.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": movie.ID}, bson.M{"$set": movie})
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *repository) DeleteMovie(ID string) error {
	movieID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": movieID})
	if err != nil {
		return err
	}	
	return nil
}
