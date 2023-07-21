package thumbHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/handle/storages/SdkApiUtils/alistApi"
	"github.com/g-mero/gotutu/utils/config"
	"path"
)

type ThumbAlist struct {
	Token     string
	Host      string
	CachePath string
}

var apiAlist alistApi.AlistApi

func (that ThumbAlist) Init() ThumbStorageMethod {
	token := config.Get("alist", "token", "YourTokenHere")
	host := config.Get("alist", "host", "http://127.0.0.1:5244")

	apiAlist = alistApi.AlistApi{
		Token: token,
		Host:  host,
	}

	return ThumbAlist{
		Token:     token,
		Host:      host,
		CachePath: config.Get("alist", "cache_path", "pics/cache_thumb"),
	}
}

func (that ThumbAlist) SaveThumbnail(img *imgHandle.ImageG, webPath string) error {
	var (
		thumbDir  string
		err       error
		thumbnail *imgHandle.ImageG
	)

	thumbDir = path.Clean(that.CachePath + "/" + path.Dir(webPath))

	thumbnail, err = img.MakeThumbnail()

	if err != nil {
		return err
	}

	err = apiAlist.UploadImg(thumbDir, thumbnail)
	if err != nil {
		return err
	}

	return nil
}

func (that ThumbAlist) GetThumbnail(webPath string) (storages.ImageInfo, error) {
	var (
		err       error
		thumbPath string
		imgInfo   storages.ImageInfo
	)

	thumbPath = path.Clean(that.CachePath + "/" + imgHandle.ThumbnailName(webPath) + ".webp")

	alistImgInfo, err := apiAlist.GetImgInfo(thumbPath)
	if err != nil {
		return imgInfo, err
	}

	imgInfo.IsLocal = false
	imgInfo.Path = alistImgInfo.RawUrl

	return imgInfo, nil
}

func (that ThumbAlist) IsThumbPath(webPath string) bool {
	webPath = path.Clean("/" + webPath)
	match, err := path.Match("/"+that.CachePath+"/", webPath)
	if err != nil {
		return true
	}

	return match
}
