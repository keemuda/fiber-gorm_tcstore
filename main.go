package main

import (
	"fmt"

	"tcstorego/router"

	"tcstorego/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("hello world!")
	app := fiber.New()
	app.Use(cors.New())
	router.SetRoutes(app)
	database.ConnectDB()

	app.Listen(":3000")
}

/*ref
-https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34
-https://github.com/gofiber/recipes/tree/master/gorm-mysql
*/
