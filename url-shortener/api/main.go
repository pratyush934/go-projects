package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pratyush934/go-projects/api/routes"
	"log"
	"os"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading  file", err.Error())
	}

	app := fiber.New()

	app.Use(logger.New())

	setUpRoutes(app)

	err = app.Listen(os.Getenv("APP_PORT"))
	if err != nil {
		return

	}
}
