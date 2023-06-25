package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dahlansrt/gofiber-crud-api/api/routes"
	"github.com/dahlansrt/gofiber-crud-api/pkg/book"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnection(mongoURI string) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database("fiber_test")

	return db, nil, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	db, cancel, err := DbConnection(mongoURI)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to the CRUD API")
	})

	api := app.Group("/api")
	routes.BookRouter(api, bookService)

	defer cancel()
	log.Fatal(app.Listen(":3000"))
}
