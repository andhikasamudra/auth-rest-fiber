package main

import (
	cg "github.com/andhikasamudra/auth-rest-fiber/config"
	"github.com/andhikasamudra/auth-rest-fiber/database"
	"github.com/andhikasamudra/auth-rest-fiber/handlers"
	"github.com/andhikasamudra/auth-rest-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := database.New()

	h := handlers.New(db.Conn, cg.NewConfig())

	routes.SetupRoutes(app, h)
	app.Listen(":3000")

	defer db.Conn.Close()
}
