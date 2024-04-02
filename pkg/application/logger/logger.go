package logger

// Logger is an interface for logging messages. It provides a standardized way to
// log messages, allowing different logging implementations to be used interchangeably.
type LoggerInterface interface {
	// Print logs a message at the given level.
	Log(level Level, args ...interface{})

	// Printf logs a formatted message at the given level.
	Logf(level Level, format string, args ...interface{})
}

// Level represents the severity level of a log message.
type Level int

const (
	// LevelDebug represents the debug level log messages.
	LevelDebug Level = iota

	// LevelInfo represents the informational level log messages.
	LevelInfo

	// LevelWarn represents the warning level log messages.
	LevelWarn

	// LevelError represents the error level log messages.
	LevelError

	// LevelFatal represents the fatal level log messages.
	LevelFatal
)

// String returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
