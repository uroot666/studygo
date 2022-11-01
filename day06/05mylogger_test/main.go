package main

import (
	"time"

	"github.com/uroot666/studygo/mylogger"
)

var log mylogger.Logger // 声明一个全局的接口变量

func main() {
	// log = mylogger.Newlog("debug")  // 终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "access.log", 1024) // 文件日志实例
	for {
		log.Debug("Debug: %v:%v", "a", "b")
		log.Trace("Trace")
		log.Info("Info")
		log.Warning("Warning")
		log.Error("Error")
		log.Fatal("Fatal")
		time.Sleep(1 * time.Second)
	}
}
