package routes

import (
	"github.com/JPSBarbosa/GO-API-REST/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupSessionRoutes(app *fiber.App) {
	app.Get("/Session/:id", controllers.GetSession)
	app.Get("/", controllers.Homepage)
	app.Get("/Teste", controllers.Teste)
}
