package hm_log

type Log interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
	Init()
}
