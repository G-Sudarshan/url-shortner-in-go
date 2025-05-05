package main

import (
	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load(".env")
	app := fiber.New()

}
