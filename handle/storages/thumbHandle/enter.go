package thumbHandle

import (
	"errors"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/utils/config"
)

type ThumbStorageMethod interface {
	// Init 初始化容器
	Init() ThumbStorageMethod
	// SaveThumbnail 保存缩略图
	SaveThumbnail(img *imgHandle.ImageG, webPath string) error
	// GetThumbnail 获取缩略图
	// webPath 网络请求的路径 eg: 202307/3424235236.png
	GetThumbnail(webPath string) (storages.ImageInfo, error)
	// IsThumbPath 判读webPath是否是缩略图
	IsThumbPath(webPath string) bool
}

// ThumbStorageRegister 缩略图容器注册
var ThumbStorageRegister = map[string]ThumbStorageMethod{
	"local":       ThumbLocal{},
	"local_cache": ThumbLocalCache{},
	"alist":       ThumbAlist{},
	"upyun":       UpyunThumb{},
}

var ErrorThumbNotExist = errors.New("缩略图文件未找到")

func ErrorIsThumbNotExist(err error) bool {
	return err == ErrorThumbNotExist
}

// InitThumbStorage 初始化原始图容器
func InitThumbStorage() ThumbStorageMethod {
	var (
		storage ThumbStorageMethod
		ok      bool
	)

	storage, ok = ThumbStorageRegister[config.Server.ThumbStorage]

	if !ok {
		storage = ThumbLocal{}.Init()
	} else {
		storage = storage.Init()
	}

	return storage
}
