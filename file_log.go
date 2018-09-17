package hm_log

import (
	"fmt"
	"os"
	"strconv"
)

type FileLog struct {
	level       int
	logPath     string
	logName     string
	file        *os.File
	warnFile    *os.File
	logDataChan chan *LogData
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

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	chanSize, e := strconv.Atoi(logChanSize)
	if e != nil {
		chanSize = 50000
	}
	logFile = &FileLog{
		level:       level,
		logPath:     logPath,
		logName:     logName,
		logDataChan: make(chan *LogData, chanSize),
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
	if f.level > DebugLevel {
		return
	}
	logData := MsgInfo(DebugLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level > TraceLevel {
		return
	}
	logData := MsgInfo(TraceLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level > InfoLevel {
		return
	}
	logData := MsgInfo(InfoLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level > WarnLevel {
		return
	}
	logData := MsgInfo(WarnLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level > ErrorLevel {
		return
	}
	logData := MsgInfo(ErrorLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level > FatalLevel {
		return
	}
	logData := MsgInfo(FatalLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Close() {
	f.file.Close()
	f.warnFile.Close()
}
