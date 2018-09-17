package hm_log

import "fmt"

var log Log

func InitLog(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		log, err = NewFileLog(config)
	case "console":
		log, err = NewConsoleLog()
	default:
		err = fmt.Errorf("unspport log name:%s", name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}

func Close() {
	log.Close()
}
