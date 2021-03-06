package main

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "/")

	log.Fatal(app.Listen(":8080"))
}
