package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	Warning = 2000

	FileNotImage = 6001
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	Warning: "警告：请求成功但是出现了一些不影响主要功能的错误",

	FileNotImage: "文件不是符合要求的图片（png, jpg, gif, webp）",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
