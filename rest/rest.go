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

	test := app.Group("test")

	test.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello test")
	})
	test.Get("/json", func(c fiber.Ctx) error {
		response := map[string]interface{}{
			"ip":        c.IP(),
			"localAddr": c.Response().LocalAddr(),
			"port":      c.Port(),
		}
		return c.JSON(response)
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	//? Define a route for the GET method on the root path '/ip'
	app.Get("/ip", func(c fiber.Ctx) error {
		//? Send a JSON response to the client by passing a map to the c.JSON function
		//? This method of creating a map allows you to initialize it with values
		response := map[string]string{
			"ip": c.IP(),
		}
		response["test"] = c.BaseURL()
		return c.JSON(response)
	})

	//? Define for the GET method on the root path '/json'
	app.Get("/json", func(c fiber.Ctx) error {
		//? pass a struct to the c.JSON function
		return c.JSON(config)
	})

	//? Define for the GET method on the root path '/json'
	app.Get("/json/:test", func(c fiber.Ctx) error {

		//? This method of creating a map allows you to add values after it is created
		data := make(map[string]interface{})
		data["name"] = c.Params("test")
		data["age"] = 30
		data["body"] = c.Body()

		//? pass a struct to the c.JSON function
		return c.JSON(data)
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
