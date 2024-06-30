package router

import (
	"tcstorego/controller"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", controller.HelloGo)
	app.Get("/dudu", controller.Dududumdum)
	app.Get("/findtc", controller.Findtc)
	app.Post("/addtc", controller.Addtc)
	app.Delete("/deletetc", controller.Deletetc)
}
