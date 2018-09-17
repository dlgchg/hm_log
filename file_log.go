package hm_log

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLog struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func NewFileLog(config map[string]string) (logFile Log, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := GetLevel(logLevel)
	logFile = &FileLog{
		level:   level,
		logPath: logPath,
		logName: logName,
	}

	logFile.Init()
	return
}

func (f *FileLog) Init() {
	filePath := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filePath, err))
	}

	f.file = file

	filePath = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filePath, err))
	}

	f.warnFile = file
}

func (f *FileLog) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		f.level = DebugLevel
	}
	f.level = level
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.level != DebugLevel {
		return
	}
	msgInfo(f.file, f.level, format, args)
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level != TraceLevel {
		return
	}
	msgInfo(f.file, f.level, format, args)
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level != InfoLevel {
		return
	}
	msgInfo(f.file, f.level, format, args)
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level != WarnLevel {
		return
	}
	msgInfo(f.warnFile, f.level, format, args)
}

func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level != ErrorLevel {
		return
	}
	msgInfo(f.warnFile, f.level, format, args)
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level != FatalLevel {
		return
	}
	msgInfo(f.warnFile, f.level, format, args)
}

func (f *FileLog) Close() {
	f.file.Close()
	f.warnFile.Close()
}

func msgInfo(file *os.File, level int, format string, args ...interface{}) {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := LogLevelString(level)
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	fmt.Fprintf(file, "%s %s [%s/%s:%d] %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}
