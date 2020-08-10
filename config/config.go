package config

import (
	"fmt"
	"new_gin_project/utils"
	"sync"
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

var s sync.Once

func Opts() Config {
	s.Do(func() {
		utils.InitTomlConfigs([]*utils.ConfigMap{
			{
				FilePath: "./conf/config.toml",
				Pointer:  &Optional,
			},
		})
	})
	return Optional
}

func InitConfigs() {

	fmt.Println(Opts())
}
