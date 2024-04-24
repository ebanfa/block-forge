package mocks

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
	loggedMessages []string
	LastMessage    string // LastMessage stores the last logged message
}

// Print logs a message at the given level.
func (m *MockLogger) Log(level logger.Level, args ...interface{}) {
	m.loggedMessages = append(m.loggedMessages, fmt.Sprint(args...))
}

// Printf logs a formatted message at the given level.
func (m *MockLogger) Logf(level logger.Level, format string, args ...interface{}) {
	m.LastMessage = fmt.Sprintf(format, args...)
	m.loggedMessages = append(m.loggedMessages, fmt.Sprint(args...))
}
