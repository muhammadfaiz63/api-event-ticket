package routes

import (
	"api-event-ticket/controllers"

	"github.com/gofiber/fiber/v2"
)

func TicketRoute(app *fiber.App) {
	app.Post("/ticket", controllers.CreateTicket)
	app.Get("/ticket/:_id", controllers.GetATicket)
	app.Put("/ticket/:_id", controllers.EditATicket)
	app.Delete("/ticket/:_id", controllers.DeleteATicket)
	app.Get("/ticket", controllers.GetAllTicket)
}
