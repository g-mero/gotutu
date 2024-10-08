// Package thumbHandle 又拍云的缩略图容器
package thumbHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/handle/storages/SdkApiUtils/upyunSDK"
	"github.com/g-mero/gotutu/utils/cache"
	"github.com/g-mero/gotutu/utils/config"
	"log"
	"path"
)

type UpyunThumb struct {
	CustomHost string // 你的加速域名加协议，例如 https://yp.test.com
	Bucket     string
	OpName     string
	Password   string
	ThumbPath  string // 缩略图保存的路径
}

var api upyunSDK.UpyunG

func (that UpyunThumb) Init() ThumbStorageMethod {
	customHost := config.Get("upyun", "customHost", "https://www.google.com")
	bucket := config.Get("upyun", "bucket", "yourBucket")
	opName := config.Get("upyun", "op_name", "operatorName")
	password := config.Get("upyun", "password", "operatorPassword")
	thumbPath := config.Get("upyun", "thumb_path", "thumbnail_cache")

	api = upyunSDK.New(bucket, opName, password)

	return UpyunThumb{
		CustomHost: customHost,
		Bucket:     bucket,
		OpName:     opName,
		Password:   password,
		ThumbPath:  thumbPath,
	}
}

func (that UpyunThumb) SaveThumbnail(img *imgHandle.ImageG, webPath string) (err error) {
	remoteDir := path.Clean(that.ThumbPath + "/" + path.Dir(webPath))

	var tinyImg *imgHandle.ImageG

	tinyImg, err = img.MakeThumbnail()
	if err != nil {
		return
	}

	err = api.UploadImg(remoteDir, tinyImg)

	return
}

func (that UpyunThumb) GetThumbnail(webPath string) (imgInfo storages.ImageInfo, err error) {
	remotePath := path.Clean("/" + that.ThumbPath + "/" + imgHandle.ThumbnailName(webPath) + ".webp")

	if !cache.Has("upyunCache"+remotePath) && !api.IsFileExist(remotePath) {
		cache.Set("upyunCache"+remotePath, nil)
		log.Println("[upyun]缩略图 " + remotePath + " 不存在(此检查对每张不存在的缩略图只进行一次，重启后才会继续检查)")
		err = ErrorThumbNotExist
		return
	}

	imgInfo.IsLocal = false
	imgInfo.Path = that.CustomHost + remotePath

	return
}

func (that UpyunThumb) IsThumbPath(webPath string) bool {
	webPath = path.Clean("/" + webPath)
	match, err := path.Match("/"+that.ThumbPath+"/", webPath)
	if err != nil {
		return true
	}

	return match
}
