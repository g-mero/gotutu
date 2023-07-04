package imgHandle

import (
	"errors"
	"github.com/g-mero/gotutu/utils/imgEncoder"
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
	Buf       []byte // 图片的字节流数据
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
	that.Buf, err = io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	that.FileName = file.Filename[:len(file.Filename)-len(path2.Ext(file.Filename))]

	return OpenFromBuffer(that.Buf, that.FileName, compress...)
}

func OpenFromBuffer(buf []byte, filename string, compress ...bool) (*ImageG, error) {
	var that = new(ImageG)

	that.Buf = buf
	that.FileName = filename

	encoder, err := imgEncoder.LoadImgFromBuffer(buf)

	if err != nil {
		return nil, err
	}

	switch encoder.ImgType {
	case imgEncoder.ImgTypeWEBP:
		that.ImageType = Webp
	case imgEncoder.ImgTypeGIF:
		that.ImageType = Gif
	case imgEncoder.ImgTypeJPEG:
		that.ImageType = Jpeg
	case imgEncoder.ImgTypePng:
		that.ImageType = Png
	default:
		return nil, errors.New("不支持的图片格式")
	}

	if len(compress) > 0 {
		if compress[0] == true {
			that.Buf, err = encoder.Compress()
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
	webp, err := imgEncoder.EncodeWebp(that.Buf, 10, 700, 800)
	if err != nil {
		return nil, err
	}

	thumbImg.Buf = webp
	thumbImg.FileName = that.FileName + "_small"
	thumbImg.ImageType = Webp

	return thumbImg, nil
}

func ThumbnailFullName(originPath string) string {
	return originPath[:len(originPath)-len(path2.Ext(originPath))] + "_small.webp"
}
