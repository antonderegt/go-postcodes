package main

import (
	"github.com/antonderegt/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/address", book.ReturnAddress)
	app.Get("/api/postcode", book.ConsumeAPI)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
