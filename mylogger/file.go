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
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

var (
	// MaxSize 日志通道缓冲区大小
	MaxSize = 50000
)

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
		logChan:     make(chan *logMsg, MaxSize),
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
	// 开启一个后台的goroutine去往文件写日志
	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件

	// 2. 备份一下 rename xx.log -> xx.log.bak202001141223
	nowStr := time.Now().Format("20060102151406000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取备份文件信息失败: %v", err)
	}

	logName := path.Join(f.filePath, fileInfo.Name())                             // 拿到当前日志完整路径
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, fileInfo.Name(), nowStr) // 备份文件路径

	// 1. 关闭当前的日志文件
	file.Close()

	os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开新的日志文件失败: %v", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.funcName, logTmp.fileName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile2, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile2
				}
				// 如果要记录的日子大于等于error级别，还要在err日志文件中再记录一遍
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.Enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getinfo(3)
		// 先把日志发送到通道中
		// 1. 造一个logMsg对象
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			// 把日志丢掉保证不出现阻塞
		}

	}
}

func (f *FileLogger) Enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("checksize 获取文件信息失败: %v\n", err)
		return false
	}
	// 如果当前文件大小大于等于日志文件的最大值就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) Debug(msg string, a ...interface{}) {
	f.log(DEBUG, msg, a...)

}

func (f *FileLogger) Trace(msg string, a ...interface{}) {
	f.log(TRACE, msg, a...)
}

func (f *FileLogger) Info(msg string, a ...interface{}) {
	f.log(INFO, msg, a...)
}

func (f *FileLogger) Warning(msg string, a ...interface{}) {
	f.log(WARNING, msg, a...)
}

func (f *FileLogger) Error(msg string, a ...interface{}) {
	f.log(ERROR, msg, a...)
}

func (f *FileLogger) Fatal(msg string, a ...interface{}) {
	f.log(FATAL, msg, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
