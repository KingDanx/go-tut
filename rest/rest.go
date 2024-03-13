package rest

import (
	"test/pg"

	"github.com/gofiber/fiber/v3"
)

func RestServer() {
	config, err := pg.GetConfig()
	if err != nil {
		panic(1)
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})
	// Define a route for the GET method on the root path '/'
	app.Get("/ip", func(c fiber.Ctx) error {
		// Send a string response to the client
		response := map[string]string{
			"ip": c.IP(),
		}
		return c.JSON(response)
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/json", func(c fiber.Ctx) error {
		return c.JSON(config)
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
