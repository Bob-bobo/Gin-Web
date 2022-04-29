package logging

import (
	logs "github.com/alecthomas/log4go"
)

var logger1File = "D:/temp/demo1.log"
var Logger1 = logs.Logger{}

var logger2File = "D:/temp/demo2.log"
var Logger2 = logs.Logger{}

func InitLog() {
	rotateStatus := false
	logs.NewConsoleLogWriter()
	log := logs.NewConsoleLogWriter()
	//可以定义输出格式
	//log.SetFormat("")
	Logger1.AddFilter("stdout", logs.INFO, log)
	Logger1.AddFilter("file", logs.DEBUG, logs.NewFileLogWriter(logger1File, rotateStatus))

	Logger2.AddFilter("stdout", logs.INFO, log)
	Logger2.AddFilter("file", logs.DEBUG, logs.NewFileLogWriter(logger2File, rotateStatus))
}
