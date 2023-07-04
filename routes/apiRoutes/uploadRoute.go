package apiRoutes

import (
	"github.com/g-mero/gotutu/api"
	"github.com/gofiber/fiber/v2"
)

type UploadRoute struct {
}

func (r UploadRoute) Init(Router fiber.Router) {
	upRouter := Router.Group("upload")

	upApi := api.GroupApp.UploadApi

	upRouter.Post("", upApi.Upload)
}
