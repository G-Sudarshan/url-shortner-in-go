package main

import (
	"fmt"
	"log"
	"os"

	"github.com/g-sudarshan/url-shortner-in-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load(".env")
	fmt.Println("Error loading .env file", err)
	fmt.Println("Error loading .env file", err)
	app := fiber.New()

	app.Use(logger.New())
	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
