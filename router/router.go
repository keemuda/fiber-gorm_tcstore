package router

import (
	"tcstorego/controller"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/Testcases", controller.FindTestcase)
	app.Post("/Testcases", controller.AddTestcase)
	app.Delete("/Testcases", controller.DeleteTestcase)
	app.Get("/download", controller.GetTestcaseFile)
}
