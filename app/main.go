package main

import (
	"time"

	"github.com/antonderegt/postcode"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func setupRoutes(app *fiber.App) {
	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/api/address", postcode.ReturnAddress)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)
	app.Get("/api/latlon", postcode.GetLatLon)
	app.Get("/api/postcode", postcode.GetPostcode)
}

func login(c *fiber.Ctx) {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t})
}

func restricted(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	c.Send("Welcome " + name)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
