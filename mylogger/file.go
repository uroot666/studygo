package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile() // 按照文件路径和名称将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileOjb, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开日志文件失败: %v", err)
		return err
	}
	errfileOjb, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开错误日志文件失败: %v", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileOjb
	f.errFileObj = errfileOjb
	return nil
}

func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.Enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getinfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			// 如果要记录的日子大于等于error级别，还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

func (f FileLogger) Enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f FileLogger) Debug(msg string, a ...interface{}) {
	f.log(DEBUG, msg, a...)

}

func (f FileLogger) Trace(msg string, a ...interface{}) {
	f.log(TRACE, msg, a...)
}

func (f FileLogger) Info(msg string, a ...interface{}) {
	f.log(INFO, msg, a...)
}

func (f FileLogger) Warning(msg string, a ...interface{}) {
	f.log(WARNING, msg, a...)
}

func (f FileLogger) Error(msg string, a ...interface{}) {
	f.log(ERROR, msg, a...)
}

func (f FileLogger) Fatal(msg string, a ...interface{}) {
	f.log(FATAL, msg, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
