package imgEncoder

import (
	"errors"
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"log"
	"math"
	"runtime"
)

var (
	boolFalse   vips.BoolParameter
	intMinusOne vips.IntParameter
)

var (
	ErrorSupportImage = errors.New("不支持的图片格式")
)

func init() {
	vips.Startup(&vips.Config{
		ConcurrencyLevel: runtime.NumCPU(),
	})

	boolFalse.Set(false)
	intMinusOne.Set(-1)
}

func imageIgnore(imageFormat vips.ImageType) bool {
	// Ignore Unknown, WebP, AVIF
	ignoreList := []vips.ImageType{vips.ImageTypeUnknown, vips.ImageTypeWEBP, vips.ImageTypeAVIF}
	for _, ignore := range ignoreList {
		if imageFormat == ignore {
			// Return err to render original image
			return true
		}
	}
	return false
}

// 图片大小缩小
func thumbNail(img *vips.ImageRef, desW, desH int) error {
	var (
		srcW = float64(img.Width())
		srcH = float64(img.PageHeight())
	)

	// 不缩放
	if img.Width() <= desW && img.PageHeight() <= desH {
		return nil
	}

	// 宽高比
	ratio := srcW / srcH

	if desW > 0 && desH <= 0 {
		return img.Thumbnail(desW, int(srcH/ratio), 0)
	}

	if desW <= 0 && desH > 0 {
		return img.Thumbnail(int(float64(desH)*ratio), desH*img.Pages(), 0)
	}

	// 自适应计算，确保宽高满足条件
	if desW > 0 && desH > 0 {
		thumbRatio := math.Min(float64(desW)/srcW, float64(desH)/srcH)

		h := int(math.Ceil(srcH * thumbRatio))
		w := int(float64(h) * ratio)

		return img.Thumbnail(w, h*img.Pages(), 0)
	}

	return nil
}

func LoadImgFromBuffer(buffer []byte) (*Encoder, error) {
	encode := new(Encoder)
	img, err := vips.LoadImageFromBuffer(buffer, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})

	if err != nil {
		return nil, err
	}
	switch img.Format() {
	case vips.ImageTypeWEBP:
		encode.ImgType = ImgTypeWEBP
	case vips.ImageTypeGIF:
		encode.ImgType = ImgTypeGIF
	case vips.ImageTypeJPEG:
		encode.ImgType = ImgTypeJPEG
	case vips.ImageTypePNG:
		encode.ImgType = ImgTypePng
	default:
		return nil, errors.New("不支持的图片格式")
	}

	encode.img = img

	return encode, nil
}

// EncodeWebp 将图片转成webp格式
// 注意 quality 1-100 , >= 100 代表无损压缩
// widthHeight 可以指定宽高，来生成thumbnail
// eg EncodeWebp(buf, 100, 250, 250)
// 这会生成一张最大宽高为250的图片，自适应的对宽高进行调整
// 你可以设置250, 0 或者 0, 250 来仅限制宽或高
func EncodeWebp(buffer []byte, quality int, widthHeight ...int) (resBuf []byte, err error) {
	var (
		width  int = 0
		height int = 0
		img    *vips.ImageRef
	)

	// 从参数中获取宽或高
	for k, v := range widthHeight {
		if k == 0 {
			width = v
		} else if k == 1 {
			height = v
		} else {
			break
		}
	}

	img, err = vips.LoadImageFromBuffer(buffer, &vips.ImportParams{
		FailOnError: boolFalse,
		NumPages:    intMinusOne,
	})

	if err != nil {
		return nil, err
	}

	if imageIgnore(img.Format()) {
		return nil, ErrorSupportImage
	}

	// 生成thumbnail
	if !(width <= 0 && height <= 0) {
		log.Println(fmt.Sprintf("[MakeThumbnail]srcWidth: %d, srcHeight: %d, pages: %d, pageHeight: %d, width: %d, height: %d",
			img.Width(), img.Height(), img.Pages(), img.PageHeight(), width, height))
		err := thumbNail(img, width, height)

		if err != nil {
			return nil, err
		}
	}

	// If quality >= 100, we use lossless mode
	if quality >= 100 {
		resBuf, _, err = img.ExportWebp(&vips.WebpExportParams{
			Lossless:        true,
			StripMetadata:   true,
			ReductionEffort: 4,
		})
	} else {
		resBuf, _, err = img.ExportWebp(&vips.WebpExportParams{
			Quality:         quality,
			Lossless:        false,
			StripMetadata:   true,
			ReductionEffort: 4,
		})
	}

	if err != nil {
		return nil, err
	}

	img.Close()

	return resBuf, nil
}
