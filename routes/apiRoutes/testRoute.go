package apiRoutes

import (
	"github.com/g-mero/gotutu/api"
	"github.com/gofiber/fiber/v2"
)

type TestRoute struct {
}

func (r TestRoute) Init(Router fiber.Router) {
	testRouter := Router.Group("test")

	testApi := api.GroupApp.TestApi

	testRouter.Get("", testApi.Test)
}
