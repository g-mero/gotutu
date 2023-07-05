package api

import (
	"github.com/g-mero/gotutu/handle"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/utils/fiberResp"
	"github.com/gofiber/fiber/v2"
	"log"
	"mime/multipart"
)

type UploadApi struct {
}

func (a UploadApi) Upload(c *fiber.Ctx) error {
	var (
		err           error
		file          *multipart.FileHeader
		img           *imgHandle.ImageG
		originStorage = handle.OriginStorage
		thumbStorage  = handle.ThumbStorage
		compress      bool
	)

	file, err = c.FormFile("image")

	// 是否压缩原图
	compress = c.QueryBool("compress")

	if err != nil {
		log.Println("[upload]get form file", err)
		return fiberResp.Error(c)
	}

	img, err = imgHandle.Open(file, compress)

	if err != nil {
		return fiberResp.ErrorMsg(c, "文件打开出错："+err.Error())
	}

	var imgUrl storages.ImageUrl

	imgUrl, err = originStorage.SaveImg(img)

	if err != nil {
		return fiberResp.ErrorMsg(c, "保存失败 "+err.Error())
	}

	if err = thumbStorage.SaveThumbnail(img, imgUrl.Path); err != nil {
		log.Println("[SaveThumbnail]保存缩略图出错： ", err)
		return fiberResp.Warn(c, "图片上传成功但是缩略图保存出现了问题："+err.Error(), imgUrl)
	}

	return fiberResp.Ok(c, imgUrl)
}
