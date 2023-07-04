// Package originHandle
// 这里是原始图片存储的处理，用于存储和处理原始图片的存放位置或形式
package originHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/utils/config"
)

type OriginStorageMethod interface {
	// Init 初始化容器
	Init() OriginStorageMethod
	// SaveImg 保存原始图片
	SaveImg(img *imgHandle.ImageG) (storages.ImageUrl, error)
	// GetImg 获取原始图片
	// webPath 网络请求的路径 eg: 202307/3424235236.png
	GetImg(webPath string) (storages.ImageInfo, error)
	// GetImageG 获取图片对象
	GetImageG(webPath string) (*imgHandle.ImageG, error)
}

// OriginStorageRegister 原始图片容器注册
var OriginStorageRegister = map[string]OriginStorageMethod{
	"local": OriginLocal{},
	"alist": OriginAlist{},
}

// InitOriginStorage 初始化原始图容器
func InitOriginStorage() OriginStorageMethod {
	var (
		storage OriginStorageMethod
		ok      bool
	)

	storage, ok = OriginStorageRegister[config.Server.OriginStorage]

	if !ok {
		storage = OriginLocal{}.Init()
	} else {
		storage = storage.Init()
	}

	return storage
}
