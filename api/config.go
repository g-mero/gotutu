package api

import (
	"github.com/g-mero/gotutu/handle"
	"github.com/g-mero/gotutu/utils/config"
	"github.com/g-mero/gotutu/utils/resp"
	"github.com/gofiber/fiber/v2"
)

type ConfApi struct {
}

func (a ConfApi) Reload(c *fiber.Ctx) error {
	config.Reload()
	handle.InitStorages()
	return resp.Ok(c)
}

func (a ConfApi) ResetApiToken(c *fiber.Ctx) error {
	token, err := config.ResetApiToken()

	if err != nil {
		return resp.ErrorMsg(c, "重载失败："+err.Error())
	}

	return resp.Ok(c, map[string]string{
		"token": token,
	})
}
