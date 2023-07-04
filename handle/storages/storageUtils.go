package storages

import (
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/utils/config"
	"net/url"
	path2 "path"
	"strings"
)

// StorageMethod 存储的基本接口
type StorageMethod interface {
	Init() StorageMethod
	SaveImg(img *imgHandle.ImageG) (ImageUrl, error)
	GetImg(path string) (ImageInfo, error)
	GetThumbnail(path string) (ImageInfo, error)
}

type ImageUrl struct {
	Url      string `json:"url"`
	ThumbUrl string `json:"thumb_url"`
	Path     string `json:"path"`
}

// ImageInfo 提供给路由api的图片信息，如果是本地图片请提供本地路径或者buffer
// 使用buffer的话，一定要提供ContentType
// 如果是远程链接，请提供path，我们会302跳转到该path所指向的url
type ImageInfo struct {
	IsLocal     bool
	Path        string
	Buf         []byte
	ContentType string
}

// MakeImgUrl path： 基础路径 eg 20230627/test.jpg
func MakeImgUrl(path string) (imgUrl ImageUrl) {
	joinPath, _ := url.JoinPath(config.Server.Host, "pic", path)
	imgUrl.Url, _ = netutil.EncodeUrl(joinPath)
	imgUrl.ThumbUrl = imgUrl.Url + "?size=small"
	imgUrl.Path = path2.Clean(path)

	return imgUrl
}

// IsImgPath 判断路径是否是图片
func IsImgPath(path string) bool {
	ext := strings.ToLower(path2.Ext(path))

	allowExt := []string{".webp", ".gif", ".png", ".jpg", "bmp"}

	return slice.Contain(allowExt, ext)
}

// GetFileNameFromPath 获取路径中图片的文件名
func GetFileNameFromPath(path string) string {
	base := path2.Base(path)

	return base[:len(base)-len(path2.Ext(base))]
}
