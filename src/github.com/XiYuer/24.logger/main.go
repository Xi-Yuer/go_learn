package main

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LoggeLevel uint16

const (
	DEBUG   LoggeLevel = 0
	Info    LoggeLevel = 1
	Warning LoggeLevel = 2
	Error   LoggeLevel = 3
)

type logger struct {
	// 日志级别
	level LoggeLevel
}

func (log *logger) Debug(s string) {
	if log.level >= DEBUG {
		printInfo(s)
	}
}
func (log *logger) Info(s string) {
	if log.level >= Info {
		printInfo(s)
	}
}
func (log *logger) Warning(s string) {
	if log.level >= Warning {
		printInfo(s)
	}
}
func (log *logger) Error(s string) {
	if log.level >= Error {
		printInfo(s)
	}
}

func getTime() string {
	return time.Now().Format("2006-01-02 03:04:05")
}

func getRuntimeInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("失败！！！")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
func printInfo(s string) {
	time := getTime()
	funcName, fileName, line := getRuntimeInfo(3)
	fmt.Printf("[%s]:[%s:%s:%d]:%s\n", time, funcName, fileName, line, s)
}

func Newlogger(level LoggeLevel) *logger {
	return &logger{
		level,
	}
}

type FileLogger struct {
	level       LoggeLevel
	filePath    string
	fileName    string
	maxFileSzie int64
}

func newFileLoger(level LoggeLevel, filePath, fileName string, maxFileSize int64) *FileLogger {
	return &FileLogger	{
		level,
		filePath,
		fileName,
		maxFileSize,
	}
}

func main() {
	log := Newlogger(Error)
	for {
		time.Sleep(time.Second * 1)
		log.Debug("调试信息")
		log.Info("普通信息")
	}
}
