package logger

import (
	"github.com/sirupsen/logrus"
)

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
