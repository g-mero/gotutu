package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length,Authorization,Content-Type",
		MaxAge:           10 * 3600,
	})
}
