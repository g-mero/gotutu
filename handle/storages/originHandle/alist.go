package originHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"github.com/g-mero/gotutu/handle/storages/SdkApiUtils/alistApi"
	"github.com/g-mero/gotutu/utils/config"
	"github.com/g-mero/gotutu/utils/request"
	"time"
)

type OriginAlist struct {
	Token string
	Host  string
	Path  string
}

var apiAlist alistApi.AlistApi

func (that OriginAlist) Init() OriginStorageMethod {
	token := config.Get("alist", "token", "YourTokenHere")
	host := config.Get("alist", "host", "http://127.0.0.1:5244")
	path := config.Get("alist", "path", "pics")
	apiAlist = alistApi.AlistApi{
		Token: token,
		Host:  host,
	}
	return OriginAlist{
		Token: token,
		Host:  host,
		Path:  path,
	}
}

func (that OriginAlist) SaveImg(img *imgHandle.ImageG) (storages.ImageUrl, error) {
	var (
		imgUrl storages.ImageUrl
		err    error
	)

	dir := time.Now().Format("20060102") + "/"

	// 保存原始图片
	err = apiAlist.UploadImg(that.Path+"/"+dir, img)
	if err != nil {
		return imgUrl, err
	}

	imgUrl = storages.MakeImgUrl(dir + img.FullName())

	return imgUrl, nil
}

func (that OriginAlist) GetImg(webPath string) (storages.ImageInfo, error) {
	var (
		err          error
		alistImgInfo alistApi.AlistApiImgInfo
		imgInfo      storages.ImageInfo
	)
	alistImgInfo, err = apiAlist.GetImgInfo(that.Path + "/" + webPath)
	if err != nil {
		return imgInfo, err
	}

	imgInfo.IsLocal = false
	imgInfo.Path = alistImgInfo.RawUrl

	return imgInfo, nil
}

func (that OriginAlist) GetImageG(webPath string) (*imgHandle.ImageG, error) {
	var (
		err    error
		imgBuf []byte
	)

	imgInfo, err := that.GetImg(webPath)
	if err != nil {
		return nil, err
	}

	resp, err := request.Get(imgInfo.Path, map[string]string{})
	if err != nil {
		return nil, err
	}

	// 获取图片的二进制code
	imgBuf = resp.Body

	return imgHandle.OpenFromBuffer(imgBuf, storages.GetFileNameFromPath(webPath))
}
