package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// Newlog 构造函数
func Newlog(levelStr string) ConsoleLogger {
	level, err := ParseLogLevel(levelStr)
	if err != nil {
		fmt.Printf("日志级别错误: %v\n", err)
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c *ConsoleLogger) Enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c *ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.Enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getinfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (c *ConsoleLogger) Debug(msg string, a ...interface{}) {
	c.log(DEBUG, msg, a...)
}

func (c *ConsoleLogger) Trace(msg string, a ...interface{}) {
	c.log(TRACE, msg, a...)
}

func (c *ConsoleLogger) Info(msg string, a ...interface{}) {
	c.log(INFO, msg, a...)
}

func (c *ConsoleLogger) Warning(msg string, a ...interface{}) {
	c.log(WARNING, msg, a...)
}

func (c *ConsoleLogger) Error(msg string, a ...interface{}) {
	c.log(ERROR, msg, a...)
}

func (c *ConsoleLogger) Fatal(msg string, a ...interface{}) {
	c.log(FATAL, msg, a...)
}
