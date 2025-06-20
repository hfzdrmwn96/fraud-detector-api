package main

import (
	"log"

	"fraud-detector-api/internal/handler"
	"fraud-detector-api/internal/middleware"
	"fraud-detector-api/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {
	pkg.InitLogger()

	app := fiber.New()

	app.Use(middleware.ErrorHandler)

	app.Post("/analyze", handler.AnalyzeTransaction)

	log.Println("Starting server at :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
