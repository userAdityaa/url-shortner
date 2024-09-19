package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	env := godotenv.Load()
	if err != nil {
		fmt.println("Error loading .env file")
	}
	app := fiber.New()
	app.Use(logger.New())
	SetUpRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
