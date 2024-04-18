package logger

import "github.com/sirupsen/logrus"

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

// LogrusLogger is a concrete implementation of LoggerInterface using Logrus.
type LogrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger creates a new instance of LogrusLogger.
func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{
		logger: logrus.New(),
	}
}

// Log logs a message at the given level.
func (l *LogrusLogger) Log(level Level, args ...interface{}) {
	switch level {
	case LevelDebug:
		l.logger.Debug(args...)
	case LevelInfo:
		l.logger.Info(args...)
	case LevelWarn:
		l.logger.Warn(args...)
	case LevelError:
		l.logger.Error(args...)
	case LevelFatal:
		l.logger.Fatal(args...)
	}
}

// Logf logs a formatted message at the given level.
func (l *LogrusLogger) Logf(level Level, format string, args ...interface{}) {
	switch level {
	case LevelDebug:
		l.logger.Debugf(format, args...)
	case LevelInfo:
		l.logger.Infof(format, args...)
	case LevelWarn:
		l.logger.Warnf(format, args...)
	case LevelError:
		l.logger.Errorf(format, args...)
	case LevelFatal:
		l.logger.Fatalf(format, args...)
	}
}
