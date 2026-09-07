package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/omerfruk/go-mongodb-api/database"
	"github.com/omerfruk/go-mongodb-api/router"
)

func main() {
	_ = godotenv.Load()

	if err := database.Setup(); err != nil {
		log.Fatalf("connect to MongoDB: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := database.Disconnect(ctx); err != nil {
			log.Printf("disconnect from MongoDB: %v", err)
		}
	}()

	app := fiber.New()
	router.Setup(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		log.Printf("stop HTTP server: %v", err)
	}
}
