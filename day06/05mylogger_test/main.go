package main

import (
	"time"

	"github.com/uroot666/studygo/mylogger"
)

func main() {
	// log := mylogger.Newlog("debug")
	log := mylogger.NewFileLogger("Info", "./", "access.log", 512)
	defer log.Close()
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
