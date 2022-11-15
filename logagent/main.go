package main

import (
	"fmt"
	"time"

	"github.com/uroot666/studygo/logagent/conf"
	"github.com/uroot666/studygo/logagent/etcd"
	"github.com/uroot666/studygo/logagent/kafka"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	// 1. 读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			// 2. 发送到kafka
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 0. 加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	// if err != nil {
	// 	fmt.Println("Load config faild")
	// }
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
	}

	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, (time.Duration(cfg.EtcdConf.Timeout) * time.Second))
	if err != nil {
		fmt.Printf("init etcd failed, err: %v", err)
	}
	fmt.Println("init etcd success")

	// 2.1 从etcd获取配置项
	logEntryConf, err := etcd.Getconf("/xxx")
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index: %v value: %v\n", index, value)
	}
	// etcd.PutTest()

	// 2. 打开日志文件准备收集日志
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed, err: %v\n", err)
	// 	return
	// }
	// fmt.Println("init taillog success")

	// run()
}
