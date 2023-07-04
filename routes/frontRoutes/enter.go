package frontRoutes

import "github.com/gofiber/fiber/v2"

func Init(r fiber.Router) {
	r.Get("/pic/+", GetImage)
}
