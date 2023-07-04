package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	FileNotImage = 6001
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	FileNotImage: "文件不是符合要求的图片（png, jpg, gif, webp）",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
