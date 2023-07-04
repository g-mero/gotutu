package frontRoutes

import (
	"github.com/g-mero/gotutu/handle"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/handle/storages/originHandle"
	"github.com/g-mero/gotutu/handle/storages/thumbHandle"
	"github.com/gofiber/fiber/v2"
)

func GetImage(c *fiber.Ctx) error {

	var (
		originStorage = handle.OriginStorage
		thumbStorage  = handle.ThumbStorage
	)

	isThumb := c.Query("size") == "small"
	webPath := c.Params("+")

	// 确保路径不会命中缩略图缓存，避免无限递归漏洞
	if thumbStorage.IsThumbPath(webPath) {
		c.Status(404)
		return c.SendString("没有找到该图片")
	}

	var imgInfo storages.ImageInfo
	var err error
	// 缩略图判断
	if isThumb {
		imgInfo, err = thumbStorage.GetThumbnail(webPath)

		if thumbHandle.ErrorIsThumbNotExist(err) {
			imgInfo, err = remakeThumb(originStorage, thumbStorage, webPath)
		}
	} else {
		imgInfo, err = originStorage.GetImg(webPath)
	}

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString(err.Error())
	}

	if imgInfo.IsLocal {
		if imgInfo.Buf != nil {
			c.Set("Content-Type", imgInfo.ContentType)
			c.Set("content-disposition", "inline")
			return c.Send(imgInfo.Buf)
		}
		return c.SendFile(imgInfo.Path)
	}

	return c.Redirect(imgInfo.Path, 302)
}

func remakeThumb(originStorage originHandle.OriginStorageMethod, thumbStorage thumbHandle.ThumbStorageMethod, webPath string) (storages.ImageInfo, error) {
	var (
		imgInfo   storages.ImageInfo
		originImg *imgHandle.ImageG
		err       error
	)

	originImg, err = originStorage.GetImageG(webPath)
	if err != nil {
		return imgInfo, err
	}

	err = thumbStorage.SaveThumbnail(originImg, webPath)
	if err != nil {
		return imgInfo, err
	}

	return thumbStorage.GetThumbnail(webPath)
}
