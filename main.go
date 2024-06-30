package main

import (
	"fmt"

	"tcstorego/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("hello world!")
	app := fiber.New()
	app.Use(cors.New())
	router.SetRoutes(app)

	app.Listen(":3000")
}

//https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34
