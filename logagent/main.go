package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/uroot666/studygo/logagent/conf"
	"github.com/uroot666/studygo/logagent/etcd"
	"github.com/uroot666/studygo/logagent/kafka"
	"github.com/uroot666/studygo/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
	wg  sync.WaitGroup
)

func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
	}

	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.MaxSize)
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
	logEntryConf, err := etcd.Getconf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}

	for index, value := range logEntryConf {
		fmt.Printf("index: %v value: %v\n", index, value)
	}

	// 3. 收集日志并发送到kafka中
	taillog.Init(logEntryConf) // 因为 newConfChan)访问了taskmgr的 newConfChan，这个chann是在taillog.Init(logEntryConf)执行的初始化
	// 2.2 派一个哨兵去监视日志收集项的变化,有变化及时通知我的logagent实现热加载配置
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外报漏的通道
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan) // 哨兵发现变化会通知上面获取的通道
	wg.Wait()
	// etcd.PutTest()
}
