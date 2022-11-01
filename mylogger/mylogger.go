package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

// Logger 接口
type Logger interface {
	Debug(msg string, a ...interface{})
	Trace(msg string, a ...interface{})
	Info(msg string, a ...interface{})
	Warning(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "NKNOWN"
	}
}

func getinfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

func ParseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("unknown")
		return UNKNOWN, err
	}
}
