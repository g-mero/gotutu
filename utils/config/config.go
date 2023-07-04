package config

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Server ServerConf
)

type ServerConf struct {
	OriginStorage string
	ThumbStorage  string
	ApiToken      string
	Host          string
}

var cfg *ini.File

var confFile = "data/config.ini"

// 初始化
func init() {
	Reload()
}

func save() {
	err := cfg.SaveTo(confFile)
	if err != nil {
		log.Println("[配置文件保存失败]", err)
		return
	}
}

func LoadServer() {
	server := cfg.Section("server")
	Server.OriginStorage = server.Key("origin_storage").MustString("local")
	Server.ThumbStorage = server.Key("thumb_storage").MustString("local")
	randApiToken, _ := getRandApiKey()
	Server.ApiToken = server.Key("api_token").MustString(randApiToken)
	Server.Host = server.Key("host").MustString("http://127.0.0.1:3095")

}

func showConf() {
	fmt.Println("[showConf]")
	defer fmt.Println("[showConf End]")
	var tmpMap map[string]interface{}

	err := cfg.MapTo(tmpMap)
	if err != nil {
		return
	}

	fmt.Println(tmpMap)
}

// ResetApiToken 重置apiToken
func ResetApiToken() (string, error) {
	apiKey, err := getRandApiKey()

	if err != nil {
		return "", err
	}

	Server.ApiToken = apiKey
	cfg.Section("server").Key("api_token").SetValue(apiKey)
	save()

	return apiKey, nil
}

func getRandApiKey() (string, error) {
	// 定义API密钥的长度（字节数）
	keyLength := 32

	// 创建一个字节数组，用于存储随机生成的数据
	key := make([]byte, keyLength)

	// 读取随机数据到字节数组中
	_, err := rand.Read(key)

	if err != nil {
		return "yourToken" + time.Now().Format("20060102"), err
	}

	// 将字节数组转换为十六进制字符串
	apiKey := hex.EncodeToString(key)

	return apiKey, err
}

func Reload() {
	var (
		err error
	)

	cfg, err = ini.Load(confFile)

	if err != nil {
		fmt.Println("没有找到配置文件！自动创建")
		cfg = ini.Empty()
		defer save()
	}

	LoadServer()

	showConf()
}

// Get 获取容器的配置项
func Get(storageName, key string, defaultValue ...string) string {
	defaultVal := ""

	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}
	storage := cfg.Section("storage." + storageName)

	ok := storage.HasKey(key)

	value := storage.Key(key).MustString(defaultVal)

	if !ok {
		save()
	}

	return value
}
