package main

import "github.com/uroot666/studygo/mylogger"

func main() {
	// log := mylogger.Newlog("debug")
	log := mylogger.NewFileLogger("Info", "./", "access.log", 10*1024*1025)
	defer log.Close()
	log.Debug("Debug: %v:%v", "a", "b")
	log.Trace("Trace")
	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
	log.Fatal("Fatal")
}
