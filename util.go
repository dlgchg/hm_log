package hm_log

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message  string
	TimeStr  string
	LevelStr string
	FileName string
	FuncName string
	LineNo   int
	IsWarn   bool
}

func GetLineInfo() (fileName, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func MsgInfo(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := LogLevelString(level)
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	isWarn := level >= WarnLevel && level <= FatalLevel

	return &LogData{
		Message:  msg,
		TimeStr:  nowStr,
		LevelStr: levelStr,
		FileName: fileName,
		FuncName: funcName,
		LineNo:   lineNo,
		IsWarn:   isWarn,
	}
}
