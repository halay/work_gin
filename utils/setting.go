package utils

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"gopkg.in/ini.v1"
	"os"
)

var (
	configFlag = flag.String("config", "", "Config filename")
	AppMode    string
	HttpPort   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisArr      string
	RedisPassword string
	RedisDb       int
	LogPath       string

	Rate     int64
	Capacity int64

	KqConf kq.KqConf
)

func init() {
	flag.Parse()
	//优化下面的代码
	if *configFlag != "" {
		fmt.Println("使用命令行参数中的配置文件")
		file, err := ini.Load(*configFlag)
		if err != nil {
			fmt.Println("配置文件读取错误，请检查文件路径:", err)
			os.Exit(1)
		}
		LoadServer(file)
		LoadData(file)
		LoadRedis(file)
		LoadRateLimit(file)
		LoadKq(file)
	} else {
		file, err := ini.Load("config/config.ini")
		if err != nil {
			fmt.Println("配置文件读取错误，请检查文件路径:", err)
			os.Exit(1)
		}
		LoadServer(file)
		LoadData(file)
		LoadRedis(file)
		LoadRateLimit(file)
		LoadKq(file)
	}
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	LogPath = file.Section("server").Key("LogPath").MustString("log")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("work_gin")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("work_gin")
}

func LoadRedis(file *ini.File) {
	RedisArr = file.Section("redis").Key("Arr").String()
	RedisPassword = file.Section("redis").Key("Password").String()
	RedisDb = file.Section("redis").Key("Db").MustInt()
}

func LoadRateLimit(file *ini.File) {
	Rate = file.Section("rate").Key("Rate").MustInt64(1)
	Capacity = file.Section("rate").Key("Capacity").MustInt64()
}
func LoadKq(file *ini.File) {
	KqConf = kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: file.Section("kq").Key("Name").MustString("kq-name"),
		},
		Brokers:   file.Section("kq").Key("Brokers").Strings(","),
		Topic:     file.Section("kq").Key("Topic").MustString("test"),
		Consumers: file.Section("kq").Key("Consumers").MustInt(8),
		Group:     file.Section("kq").Key("Group").MustString("test"),
		Offset:    file.Section("kq").Key("Offset").MustString("last"),
	}
}
