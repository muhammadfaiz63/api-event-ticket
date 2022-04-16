package routes

import (
	"api-event-ticket/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:_id", controllers.GetAUser)
	app.Put("/user/:_id", controllers.EditAUser)
	app.Delete("/user/:_id", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}
