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
		// 初始化的时候起了多少个tailtask，都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[mk] = tailObj
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
			for _, conf := range newConfg {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					// 原来就有
					continue
				} else {
					// 新增
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = tailObj
				}
			}
			// 找出原来 t.taskMap 有，但是newConf中没有的
			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConfg {
					if c2.Path == c1.Path && c2.Path == c1.Path {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.taskMap[mk].cancelFunc()
				}

			}
		default:
			time.Sleep(time.Second * 1)

		}
	}
}

// 向外报漏 taskMgr 的 newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newconfChan
}
