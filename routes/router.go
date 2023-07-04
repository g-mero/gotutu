package routes

import (
	"github.com/g-mero/gotutu/routes/apiRoutes"
	"github.com/g-mero/gotutu/routes/frontRoutes"
	"github.com/g-mero/gotutu/routes/middleware"
	"github.com/gofiber/fiber/v2"
	"log"
)

func InitRouter() {
	app := fiber.New(fiber.Config{UnescapePath: true})

	app.Use(middleware.Cors())

	frontRoutes.Init(app)

	api := app.Group("api").Use(middleware.Auth)

	apiRoutes.RoutesApp.Init(api)
	log.Fatal(app.Listen(":3095"))
}
