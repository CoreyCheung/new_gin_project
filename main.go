package main

import (
	"fmt"
	"new_gin_project/config"
	"new_gin_project/db"
	"new_gin_project/router"
	"new_gin_project/utils"

	"time"
)

func main() {

	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	db.DBInit()
	router.RouteInit()
	fmt.Println(M(10000))
}
func m(int64) int64
func M(m int64) int64 {
	return M(m)
}
