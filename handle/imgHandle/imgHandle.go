package imgHandle

import (
	"errors"
	"github.com/g-mero/goimgp"
	"io"
	"mime/multipart"
	path2 "path"
)

const (
	Jpeg int = 1
	Png  int = 2
	Gif  int = 3
	Webp int = 4
)

type ImageG struct {
	encoder   *goimgp.Encoder
	Data      []byte
	FileName  string // 图片的名称（不带后缀）
	ImageType int    // 图片的类型
}

// Open /**
func Open(file *multipart.FileHeader, compress ...bool) (*ImageG, error) {
	var that = new(ImageG)

	// 开始转码
	data, err := file.Open()

	if err != nil {
		return that, err
	}
	buf, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	that.FileName = file.Filename[:len(file.Filename)-len(path2.Ext(file.Filename))]

	return OpenFromBuffer(buf, that.FileName, compress...)
}

func OpenFromBuffer(buf []byte, filename string, compress ...bool) (*ImageG, error) {
	var that = new(ImageG)
	var err error

	that.FileName = filename
	that.encoder, err = goimgp.LoadImgFromBuffer(buf)

	if err != nil {
		return nil, err
	}

	switch that.encoder.Format {
	case goimgp.ImgTypeWEBP:
		that.ImageType = Webp
	case goimgp.ImgTypeGIF:
		that.ImageType = Gif
	case goimgp.ImgTypeJPEG:
		that.ImageType = Jpeg
	case goimgp.ImgTypePng:
		that.ImageType = Png
	default:
		return nil, errors.New("不支持的图片格式")
	}

	if len(compress) > 0 {
		if compress[0] == true {
			that.Data, err = that.encoder.LossLess()
		}
	}

	return that, err
}

func (that *ImageG) FullName() string {
	var extMap = map[int]string{
		Jpeg: ".jpg",
		Webp: ".webp",
		Gif:  ".gif",
		Png:  ".png",
	}

	return that.FileName + extMap[that.ImageType]
}

func (that *ImageG) ContentType() string {
	var extMap = map[int]string{
		Jpeg: "image/jpeg",
		Webp: "image/webp",
		Gif:  "image/gif",
		Png:  "image/png",
	}

	return extMap[that.ImageType]
}

// MakeThumbnail 生成缩略图
func (that *ImageG) MakeThumbnail() (*ImageG, error) {
	thumbImg := new(ImageG)
	webp, err := that.encoder.Tiny(700, 600)
	if err != nil {
		return nil, err
	}
	thumbImg.encoder, err = goimgp.LoadImgFromBuffer(webp)
	if err != nil {
		return nil, err
	}
	thumbImg.Data = webp
	thumbImg.FileName = that.FileName + "_small"
	thumbImg.ImageType = Webp
	return thumbImg, nil
}

// ThumbnailName 根据原始图片的路径，生成缩略图的路径，注意只有名字没有后缀
func ThumbnailName(originPath string) string {
	return originPath[:len(originPath)-len(path2.Ext(originPath))] + "_small"
}
