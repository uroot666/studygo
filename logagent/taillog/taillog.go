package taillog

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
	"github.com/uroot666/studygo/logagent/kafka"
)

// 专门从日志文件收集日志的模块

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

func (t TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err: ", err)
		return
	}
	go t.run()
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			// kafka.SendToKafka(t.topic, line.Text)
			// 先把日志数据发送到一个通道中
			// kafka 那个包中有单独的goroutine去取日志发送到kafka
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
