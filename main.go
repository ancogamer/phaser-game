package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "/")

	log.Fatal(app.Listen(os.Getenv("$PORT")))
}
