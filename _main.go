package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dahlansrt/gofiber-crud-api/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const (
	dbName = "fiber_test"
)

func Connect(mongoURI string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func DbConnect(mongoURI string) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database(dbName)

	return db, nil, nil
}

func _main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	if err = Connect(mongoURI); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []entities.Employee = make([]entities.Employee, 0)
		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})

	app.Get("/employee/:id", func(c *fiber.Ctx) error {
		params := c.Params("id")
		_id, err := primitive.ObjectIDFromHex(params)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: _id}}
		var result entities.Employee

		if err := mg.Db.Collection("employees").FindOne(c.Context(), filter).Decode(&result); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(entities.Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee.ID = ""
		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		// decode the Mongo record into Employee
		createdEmployee := &entities.Employee{}
		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)
	})

	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		employeeId, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.SendStatus(400)
		}
		employee := new(entities.Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		query := bson.D{{Key: "_id", Value: employeeId}}
		update := bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		}}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err != mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}
		employee.ID = idParam

		return c.Status(200).JSON(employee)
	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		employeeId, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}
		query := bson.D{{Key: "_id", Value: employeeId}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
		if err != nil {
			return c.SendStatus(500)
		}
		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		return c.SendStatus(204)
	})

	log.Fatal(app.Listen(":3000"))
}
