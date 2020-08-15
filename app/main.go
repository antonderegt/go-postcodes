package main

import (
	"github.com/antonderegt/postcode"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/address", postcode.ReturnAddress)
	app.Get("/api/latlon", postcode.GetLatLon)
	app.Get("/api/postcode", postcode.GetPostcode)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
