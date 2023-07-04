package apiRoutes

import "github.com/gofiber/fiber/v2"

type RouterGroup struct {
	TestRoute
	UploadRoute
	ConfRoute
}

func (g RouterGroup) Init(r fiber.Router) {
	g.TestRoute.Init(r)
	g.UploadRoute.Init(r)
	g.ConfRoute.Init(r)
}

var RoutesApp = new(RouterGroup)
