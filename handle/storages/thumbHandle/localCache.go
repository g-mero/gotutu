// Package thumbHandle 将缩略图保存到本地缓存及内存中
// 默认会申请32mb的内存作为cache(理论可存储1000张缩略图，用完之后也会自动清理最早存入的图片)
package thumbHandle

import (
	"errors"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/utils/cache"
)

type ThumbLocalCache struct {
}

func (that ThumbLocalCache) Init() ThumbStorageMethod {
	return ThumbLocalCache{}
}

func (that ThumbLocalCache) SaveThumbnail(img *imgHandle.ImageG, webPath string) error {
	thumbnail, err := img.MakeThumbnail()

	if err != nil {
		return err
	}
	cache.Set(webPath, thumbnail.Buf)

	return nil
}

func (that ThumbLocalCache) GetThumbnail(webPath string) (storages.ImageInfo, error) {
	var (
		imgInfo storages.ImageInfo
	)

	buf := cache.Get(webPath)
	if buf == nil {
		return imgInfo, errors.New("缓存中没有找到")
	}

	imgInfo.IsLocal = true
	imgInfo.Buf = buf
	imgInfo.ContentType = "image/webp"

	return imgInfo, nil
}

func (that ThumbLocalCache) IsThumbPath(webPath string) bool {
	return webPath == ""
}
