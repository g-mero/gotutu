package alistApi

import (
	"encoding/json"
	"errors"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/utils/request"
	"github.com/gofiber/fiber/v2"
	path2 "path"
)

type AlistApi struct {
	Token string
	Host  string
}

type AlistImgInfo struct {
	RawUrl string
}

func (that AlistApi) GetImgInfo(remotePath string) (AlistImgInfo, error) {
	var (
		err     error
		apiUrl  = that.Host + "/api/fs/get"
		imgInfo AlistImgInfo
	)

	resp, err := request.Post(apiUrl, map[string]string{"Authorization": that.Token}, fiber.Map{"path": remotePath})
	if err != nil {
		return imgInfo, err
	}

	var body fiber.Map
	err = json.Unmarshal(resp.Body, &body)

	if err != nil {
		return imgInfo, err
	}

	if body["code"].(float64) != 200 {
		msg := body["message"].(string)
		return imgInfo, errors.New(msg)
	}

	imgInfo.RawUrl = body["data"].(map[string]interface{})["raw_url"].(string)

	return imgInfo, nil
}

// UploadImg 上传图片到alist空间
func (that AlistApi) UploadImg(remoteDir string, img *imgHandle.ImageG) error {
	header := map[string]string{
		"Authorization": that.Token,
		"File-Path":     path2.Clean(remoteDir + "/" + img.FullName()),
	}

	apiUrl := that.Host + "/api/fs/form"
	res, err := request.Put(apiUrl, header, img.Data)
	if err != nil {
		return err
	}

	var body fiber.Map
	err = json.Unmarshal(res.Body, &body)

	if err != nil {
		return err
	}

	if body["code"].(float64) != 200 {
		return errors.New("上传失败: " + body["message"].(string))
	}

	return nil
}
