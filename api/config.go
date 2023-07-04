package api

import (
	"github.com/g-mero/gotutu/utils/config"
	"github.com/g-mero/gotutu/utils/fiberResp"
	"github.com/gofiber/fiber/v2"
)

type ConfApi struct {
}

func (a ConfApi) Reload(c *fiber.Ctx) error {
	config.Reload()
	return fiberResp.Ok(c)
}

func (a ConfApi) ResetApiToken(c *fiber.Ctx) error {
	token, err := config.ResetApiToken()

	if err != nil {
		return fiberResp.ErrorMsg(c, "重载失败："+err.Error())
	}

	return fiberResp.OkWithData(c, map[string]string{
		"token": token,
	})
}
