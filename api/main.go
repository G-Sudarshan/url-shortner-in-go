package main

import (
	"log"
	"os"

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
	if err != nil {
		fmt.println("Error loading .env file", err)
	}
	app := fiber.New()

	app.Use(logger.New())
	setUpRoutes(app)

	log.Fatal(app.Listen(os.GetEnv("APP_PORT")))
}
