package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

type mockLogger struct {
	msgs []string
}

func (l *mockLogger) Errorf(format string, args ...interface{}) {
	l.msgs = append(l.msgs, fmt.Sprintf("[ERROR] "+format, args...))
}

func (l *mockLogger) Warningf(format string, args ...interface{}) {
	l.msgs = append(l.msgs, fmt.Sprintf("[WARNING] "+format, args...))
}

func TestAssignmentChecker(t *testing.T) {
	logger := &mockLogger{}
	run([]string{filepath.Join("testdata", "enum_assignment.go")}, "Enum", logger)

	t.Logf("Messages: %v", logger.msgs)
}

func TestSwitchChecker(t *testing.T) {
	logger := &mockLogger{}
	run([]string{filepath.Join("testdata", "enum_switch.go")}, "Enum", logger)

	t.Logf("Messages: %v", logger.msgs)
}
