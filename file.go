package hm_log

import (
	"fmt"
	"os"
	"strconv"
)

type FileLog struct {
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

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	chanSize, e := strconv.Atoi(logChanSize)
	if e != nil {
		chanSize = 50000
	}
	logFile = &FileLog{
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

	go f.writeLogBackGround()
}

func (f *FileLog) writeLogBackGround() {
	for logData := range f.logDataChan {
		var file = f.file
		if logData.IsWarn {
			file = f.warnFile
		}
		fmt.Fprintf(file, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr,
			logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
	}
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	logData := MsgInfo(DebugLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	logData := MsgInfo(TraceLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Info(format string, args ...interface{}) {
	logData := MsgInfo(InfoLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	logData := MsgInfo(WarnLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Error(format string, args ...interface{}) {
	logData := MsgInfo(ErrorLevel, format, args)
	select {
	case f.logDataChan <- logData:
	default:
	}
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
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
