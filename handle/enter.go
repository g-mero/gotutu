package handle

import (
	"github.com/g-mero/gotutu/handle/storages/originHandle"
	"github.com/g-mero/gotutu/handle/storages/thumbHandle"
)

var (
	ThumbStorage  thumbHandle.ThumbStorageMethod
	OriginStorage originHandle.OriginStorageMethod
)

func InitStorages() {
	thumbHandle.InitThumbStorage()
	originHandle.InitOriginStorage()
}
