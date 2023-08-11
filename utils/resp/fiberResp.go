package resp

import (
	"github.com/g-mero/gotutu/utils/errmsg"
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, code ...int) error {
	codeHere := errmsg.ERROR

	if len(code) > 0 {
		codeHere = code[0]
	}
	return c.JSON(fiber.Map{
		"code":    codeHere,
		"message": errmsg.GetErrMsg(codeHere),
	})
}

func ErrorMsg(c *fiber.Ctx, msg string) error {
	return c.JSON(fiber.Map{
		"code":    500,
		"message": msg,
	})
}

func Ok(c *fiber.Ctx, data ...interface{}) error {
	code := errmsg.SUCCESS

	return c.JSON(fiber.Map{
		"code":    code,
		"data":    getFirstElementInSlice(data),
		"message": errmsg.GetErrMsg(code),
	})
}

func Warn(c *fiber.Ctx, msg string, data ...interface{}) error {

	return c.JSON(fiber.Map{
		"code":    errmsg.Warning,
		"data":    getFirstElementInSlice(data),
		"message": msg,
	})
}

func getFirstElementInSlice(data []interface{}) (el interface{}) {
	if len(data) > 0 {
		el = data[0]
	}
	return
}
