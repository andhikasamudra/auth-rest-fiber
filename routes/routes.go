package routes

import (
	"github.com/andhikasamudra/auth-rest-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handlers.Handler) {
	app.Get("/users", h.GetUsers)
	app.Post("/users", h.CreateUser)

	app.Post("/auth/login", h.Login)
}
