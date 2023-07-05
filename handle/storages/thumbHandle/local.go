package thumbHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"log"
	"os"
	path2 "path"
)

type ThumbLocal struct {
	CachePath string
}

func (that ThumbLocal) IsThumbPath(webPath string) bool {
	webPath = path2.Clean("/" + webPath)
	match, err := path2.Match("/"+that.CachePath+"/", webPath)
	if err != nil {
		return true
	}

	return match
}

func (that ThumbLocal) Init() ThumbStorageMethod {
	return ThumbLocal{CachePath: "data/pic/cache_thumb"}
}

func (that ThumbLocal) SaveThumbnail(img *imgHandle.ImageG, webPath string) error {
	var (
		thumbDir  string
		err       error
		thumbnail *imgHandle.ImageG
	)

	thumbDir, err = that.mkThumbDir(path2.Dir(webPath))
	if err != nil {
		return err
	}

	thumbnail, err = img.MakeThumbnail()
	if err != nil {
		return err
	}

	err = os.WriteFile(thumbDir+"/"+thumbnail.FullName(), thumbnail.Buf, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (that ThumbLocal) GetThumbnail(webPath string) (storages.ImageInfo, error) {
	var (
		err       error
		thumbPath string
		imgInfo   storages.ImageInfo
	)

	thumbPath = path2.Clean(that.CachePath + "/" + imgHandle.ThumbnailName(webPath) + ".webp")

	_, err = os.Stat(thumbPath)

	if os.IsNotExist(err) {
		return imgInfo, ErrorThumbNotExist
	}

	imgInfo.IsLocal = true
	imgInfo.Path = thumbPath

	return imgInfo, err
}

// 创建一个用于存放缩略图的目录
func (that ThumbLocal) mkThumbDir(dir string) (string, error) {
	thumbDir := path2.Clean(that.CachePath + "/" + dir)

	_, err := os.Stat(thumbDir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(thumbDir, 0755)

		if err != nil {
			log.Println("[upload]mkDir ", err)
			return "", err
		}
	}

	return thumbDir, nil
}
