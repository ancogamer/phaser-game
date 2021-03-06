package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "/")

	fmt.Println(os.Getenv("PORT"))

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
