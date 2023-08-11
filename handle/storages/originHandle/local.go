package originHandle

import (
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/g-mero/gotutu/handle/storages"
	"log"
	"os"
	path2 "path"
)

type OriginLocal struct {
	RootPath string
}

func (that OriginLocal) GetImageG(webPath string) (*imgHandle.ImageG, error) {
	info, err := that.GetImg(webPath)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(info.Path)
	if err != nil {
		return nil, err
	}

	img, err := imgHandle.OpenFromBuffer(file, storages.GetFileNameFromPath(webPath))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (that OriginLocal) Init() OriginStorageMethod {
	return OriginLocal{
		RootPath: "data/pic",
	}
}

func (that OriginLocal) SaveImg(img *imgHandle.ImageG) (storages.ImageUrl, error) {
	var (
		imgUrl storages.ImageUrl
		err    error
	)

	dateDir := storages.MakeDateDir()
	baseDir := that.mkFileDir(dateDir)

	// 保存原始图片到本地
	err = os.WriteFile(baseDir+"/"+img.FullName(), img.Data, 0644)
	if err != nil {
		return imgUrl, err
	}
	imgUrl = storages.MakeImgUrl(dateDir + "/" + img.FullName())

	return imgUrl, nil
}

func (that OriginLocal) GetImg(webPath string) (storages.ImageInfo, error) {
	var (
		imgInfo storages.ImageInfo
	)
	src := that.RootPath + "/" + webPath

	_, err := os.Stat(src)

	imgInfo.IsLocal = true
	imgInfo.Path = src

	return imgInfo, err
}

// 创建一个用于存储的目录
func (that OriginLocal) mkFileDir(path string) string {
	dir := path2.Clean(that.RootPath + "/" + path)

	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)

		if err != nil {
			log.Println("[upload]mkDir ", err)
		}
	}

	return dir
}
