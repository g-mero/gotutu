package middleware

import (
	"github.com/g-mero/gotutu/utils/config"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	checkToken := strings.SplitN(token, " ", 2)

	if len(checkToken) != 2 || checkToken[0] != "Bearer" || checkToken[1] != config.Server.ApiToken {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("无权访问")
	}

	return c.Next()
}
