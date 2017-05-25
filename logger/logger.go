package logger

import (
	"fmt"
	"io"
)

type ILogger interface {
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type Logger struct {
	writer io.Writer
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{w}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, format+"\n", args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, format+"\n", args...)
}
