package hm_log

const (
	DebugLevel = iota
	TraceLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

func LogLevelString(level int) (levelStr string) {
	switch level {
	case DebugLevel:
		levelStr = "DEBUG"
	case TraceLevel:
		levelStr = "TRACE"
	case InfoLevel:
		levelStr = "INFO"
	case WarnLevel:
		levelStr = "WARN"
	case ErrorLevel:
		levelStr = "ERROR"
	case FatalLevel:
		levelStr = "FATAL"
	}
	return
}

func GetLevel(level string) int {
	switch level {
	case "debug", "Debug", "DEBUG":
		return DebugLevel
	case "trace", "Trace", "TRACE":
		return TraceLevel
	case "info", "Info", "INFO":
		return InfoLevel
	case "warn", "Warn", "WARN":
		return WarnLevel
	case "error", "Error", "ERROR":
		return ErrorLevel
	case "fatal", "Fatal", "FATAL":
		return FatalLevel
	}
	return DebugLevel
}
