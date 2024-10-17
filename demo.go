package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodernubie/openaidemo/demo1"
	"github.com/kodernubie/openaidemo/demo2"
	"github.com/kodernubie/openaidemo/demo3"
	"github.com/kodernubie/openaidemo/demo4"
)

func main() {

	app := fiber.New()

	demo1.Init(app)
	demo2.Init(app)
	demo3.Init(app)
	demo4.Init(app)

	app.Static("/", "./web")

	app.Listen(":3000")
}
