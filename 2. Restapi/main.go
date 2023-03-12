package main

import (
	"my-rest-api/configs"
	"my-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// creating a fiber app
	app := fiber.New()

	// connecting to the db
	configs.ConnectDB()

	// connecting the routes
	routes.UserRoute(app)

	// listening on port 6000
	app.Listen(":6000")
}
