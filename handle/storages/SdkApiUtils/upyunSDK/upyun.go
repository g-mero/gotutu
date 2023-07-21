package upyunSDK

import (
	"bytes"
	"github.com/g-mero/gotutu/handle/imgHandle"
	"github.com/upyun/go-sdk/v3/upyun"
	"path"
	"strconv"
)

type UpyunG struct {
	core *upyun.UpYun
}

// New 新建一个sdk实例
func New(bucket, opName, password string) (upyunG UpyunG) {
	upyunG.core = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   bucket,
		Operator: opName,
		Password: password,
	})

	return
}

func (that UpyunG) UploadImg(remoteDir string, img *imgHandle.ImageG) error {
	var (
		err error
	)

	err = that.core.Put(&upyun.PutObjectConfig{
		Path:    path.Clean(remoteDir + "/" + img.FullName()),
		Reader:  bytes.NewReader(img.Buf),
		Headers: map[string]string{"Content-Length": strconv.Itoa(len(img.Buf))},
	})
	if err != nil {
		return err
	}

	return nil
}

// IsFileExist 判断文件是否存在
func (that UpyunG) IsFileExist(remotePath string) bool {

	_, err := that.core.GetInfo(remotePath)

	return !upyun.IsNotExist(err)
}
