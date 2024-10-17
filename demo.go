package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodernubie/openaidemo/demo1"
)

func main() {

	app := fiber.New()

	demo1.Init(app)

	app.Static("/", "./web")

	app.Listen(":3000")
}
