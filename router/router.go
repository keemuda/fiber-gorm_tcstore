package router

import (
	"tcstorego/controller"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", controller.HelloGo)
	app.Get("/dudu", controller.Dududumdum)
	app.Get("/findtc", controller.FindTestcase)
	app.Post("/addtc", controller.AddTestcase)
	app.Delete("/deletetc", controller.DeleteTestcase)
}
