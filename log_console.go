package hm_log

import "fmt"

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
	msInfo(format, args...)
}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	msInfo(format, args...)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	msInfo(format, args...)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	msInfo(format, args...)
}

func (c *ConsoleLog) Error(format string, args ...interface{}) {
	msInfo(format, args...)
}

func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	msInfo(format, args...)
}

func (c *ConsoleLog) Close() {
}

func msInfo(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}
