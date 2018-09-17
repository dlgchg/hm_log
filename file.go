package hm_log

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLog struct {
	logPath       string
	logName       string
	file          *os.File
	warnFile      *os.File
	logDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
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

	var logSplitType = LogSplitTypeSize
	var logSplitSize int64

	splitType, ok := config["log_split_type"]
	if !ok {
		splitType = "hour"
	} else {
		if splitType == "size" {
			splitSize, ok := config["log_split_size"]
			if !ok {
				splitSize = "104857600"
			}

			logSplitSize, err = strconv.ParseInt(splitSize, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}

			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}

	logFile = &FileLog{
		logPath:       logPath,
		logName:       logName,
		logDataChan:   make(chan *LogData, chanSize),
		logSplitType:  logSplitType,
		logSplitSize:  logSplitSize,
		lastSplitHour: time.Now().Hour(),
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

		f.checkSplitFile(logData.IsWarn)

		fmt.Fprintf(file, "%s %s [%s/%s:%d] %s\n", logData.TimeStr, logData.LevelStr,
			logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
	}
}

func (f *FileLog) checkSplitFile(isWarn bool) {
	if f.logSplitType == LogSplitTypeHour {
		f.splitHour(isWarn)
		return
	}
	f.splitSize(isWarn)
}

func (f *FileLog) splitSize(isWarn bool) {
	file := f.file
	if isWarn {
		file = f.warnFile
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return
	}

	fileSize := fileInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}

	var backupFileName string
	var fileName string

	now := time.Now()

	if isWarn {
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%s", f.logPath, f.logName, now.Format("20060102150405"))
		fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%s", f.logPath, f.logName, now.Format("20060102150405"))
		fileName = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	file.Close()
	os.Rename(fileName, backupFileName)

	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if isWarn {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *FileLog) splitHour(isWarn bool) {

	now := time.Now()
	hour := now.Hour()

	if hour == f.lastSplitHour {
		return
	}

	f.lastSplitHour = hour

	var backupFileName string
	var fileName string

	if isWarn {
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%s", f.logPath, f.logName, now.Format("20060102150405"))
		fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%s", f.logPath, f.logName, now.Format("20060102150405"))
		fileName = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	file := f.file
	if isWarn {
		file = f.warnFile
	}
	file.Close()
	os.Rename(fileName, backupFileName)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if isWarn {
		f.warnFile = file
	} else {
		f.file = file
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
