package main

import (
	"new_gin_project/config"
	"new_gin_project/db"
	"new_gin_project/http"
	"new_gin_project/utils"
	"os"
	"os/signal"
	"syscall"

	"new_gin_project/service"
	"time"

	"github.com/sirupsen/logrus"
)

var s *service.Service

func main() {

	config.InitConfigs() //初始化配置
	utils.InitLogrus(config.Opts().LogPath, config.Opts().LogFileName, time.Duration(24*config.Optional.LogmaxAge)*time.Hour, time.Duration(config.Optional.LogrotationTime)*time.Hour)
	db.DBInit()
	s = service.NewService()
	go http.RouteInit(s)
	signalHandler()
}
func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			logrus.Info("get a signal %s, stop the push-admin process", si.String())
			//	s.Close()
			//	s.Wait()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
