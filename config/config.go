package config

import (
	"fmt"
	"new_gin_project/utils"
)

type Config struct {
	ExtranetAddr string // 外网地址
	LocalAddress string // 内网地址
	WebPort      string // 端口

	MysqlStr     string
	DBLog        uint8
	DatabaseName string

	RedisAddr string
	RedisPwd  string
	RedisKey  string

	LogmaxAge       int64 // 日志最大保存时间（天）
	LogrotationTime int64 // 日志切割时间间隔（小时）
	LogPath         string
	LogFileName     string
}

var Optional = Config{}

func Opts() Config {
	return Optional
}

func InitConfigs() {
	utils.InitTomlConfigs([]*utils.ConfigMap{
		{
			FilePath: "./conf/config1.toml",
			Pointer:  &Optional,
		},
	})
	fmt.Println(Optional)
}