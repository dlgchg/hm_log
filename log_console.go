package hm_log

import (
	"fmt"
	"os"
)

type ConsoleLog struct {
	level int
}

func NewConsoleLog(config map[string]string) (logCon Log, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := GetLevel(logLevel)
	logCon = &ConsoleLog{
		level: level,
	}
	return
}

func (c *ConsoleLog) Init() {
}

func (c *ConsoleLog) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		c.level = DebugLevel
	}
	c.level = level
}

func (c *ConsoleLog) Debug(format string, args ...interface{}) {
	if c.level > DebugLevel {
		return
	}
	logData := MsgInfo(DebugLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	if c.level > TraceLevel {
		return
	}
	logData := MsgInfo(TraceLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	if c.level > InfoLevel {
		return
	}
	logData := MsgInfo(InfoLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	if c.level > WarnLevel {
		return
	}
	logData := MsgInfo(WarnLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Error(format string, args ...interface{}) {
	if c.level > ErrorLevel {
		return
	}
	logData := MsgInfo(ErrorLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	if c.level > FatalLevel {
		return
	}
	logData := MsgInfo(FatalLevel, format, args)
	fmt.Fprintf(os.Stdout, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLog) Close() {
}
