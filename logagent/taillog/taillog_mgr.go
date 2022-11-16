package taillog

import (
	"fmt"
	"time"

	"github.com/uroot666/studygo/logagent/etcd"
)

var taskMgr *tailLogMgr

// tailtask 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newconfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前日志收集项配置记录起来
		taskMap:     make(map[string]*TailTask, 16),
		newconfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		NewTailTask(logEntry.Path, logEntry.Topic)
	}

	go taskMgr.run()
}

// 监听自己的newConfChan, 有了新的配置过来之后就做对应的处理

func (t tailLogMgr) run() {
	for {
		select {
		case newConfg := <-t.newconfChan:
			// 1. 配置新增
			// 2. 配置删除
			// 3. 配置变更
			fmt.Println("新的配置来了,", newConfg)
		default:
			time.Sleep(time.Second * 1)

		}
	}
}

// 向外报漏 taskMgr 的 newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newconfChan
}
