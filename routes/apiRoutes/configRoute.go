package apiRoutes

import (
	"github.com/g-mero/gotutu/api"
	"github.com/gofiber/fiber/v2"
)

type ConfRoute struct {
}

func (a ConfRoute) Init(Router fiber.Router) {
	r := Router.Group("conf")

	rApi := api.GroupApp.ConfApi

	r.Post("reload", rApi.Reload)

	r.Post("resetToken", rApi.ResetApiToken)
}
