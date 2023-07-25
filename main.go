package main

import (
	"github.com/g-mero/gotutu/handle"
	"github.com/g-mero/gotutu/routes"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

// 初始化
func init() {
	// 记录核心日志
	logFile := &lumberjack.Logger{
		Filename:   "data/log/core.log",
		MaxSize:    1,     // 文件大小MB
		MaxBackups: 1,     // 最大保留日志文件数量
		MaxAge:     28,    // 保留天数
		Compress:   false, // 是否压缩
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
}

func main() {
	handle.InitStorages()
	routes.InitRouter()
}
