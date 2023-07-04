package fiberResp

import (
	"github.com/g-mero/gotutu/utils/errmsg"
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx) error {
	code := errmsg.ERROR
	return c.JSON(fiber.Map{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

func ErrorCode(c *fiber.Ctx, code int) error {
	return c.JSON(fiber.Map{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

func ErrorMsg(c *fiber.Ctx, msg string) error {
	return c.JSON(fiber.Map{
		"code":    500,
		"message": msg,
	})
}

func Ok(c *fiber.Ctx) error {
	code := errmsg.SUCCESS
	return c.JSON(fiber.Map{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

func OkWithData(c *fiber.Ctx, data interface{}) error {
	code := errmsg.SUCCESS
	return c.JSON(fiber.Map{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
