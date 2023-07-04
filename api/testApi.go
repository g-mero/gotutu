package api

import (
	"github.com/g-mero/gotutu/utils/errmsg"
	"github.com/gofiber/fiber/v2"
)

type TestApi struct {
}

func (a TestApi) Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": errmsg.GetErrMsg(200),
	})
}
