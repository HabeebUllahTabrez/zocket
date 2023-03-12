// File responsible for handling all the url endpoints and assigning them to individual handler functions

package routes

import (
	"my-rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Get("/", controllers.GetHome)

	app.Get("/students", controllers.GetAllStudents)

	app.Get("/student/:userId", controllers.GetAStudent)

	app.Post("/student", controllers.CreateStudent)

	app.Put("/student/:userId", controllers.EditAStudent)

	app.Delete("/student/:userId", controllers.DeleteAStudent)

}
