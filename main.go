package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Static("/", "./", fiber.Static{
		Compress: true,
		Browse:   true,
	})
	err := app.Listen(":80")
	if err != nil {
		log.Fatal(err)
	}
}
