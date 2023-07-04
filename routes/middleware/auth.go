package middleware

import (
	"github.com/g-mero/gotutu/utils/config"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token != config.Server.ApiToken {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("无权访问")
	}

	return c.Next()
}
