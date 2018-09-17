package hm_log

import (
	"fmt"
	"os"
)

type ConsoleLog struct{}

func NewConsoleLog() (logCon Log, err error) {
	logCon = &ConsoleLog{}
	return
}

func (c *ConsoleLog) Init() {
}

func (c *ConsoleLog) Debug(format string, args ...interface{}) {
	logData := MsgInfo(DebugLevel, format, args)
	fmt.Println(logData.FuncName)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	logData := MsgInfo(TraceLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	logData := MsgInfo(InfoLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	logData := MsgInfo(WarnLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Error(format string, args ...interface{}) {
	logData := MsgInfo(ErrorLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	logData := MsgInfo(FatalLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Close() {
}
