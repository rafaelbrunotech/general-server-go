package logger

import (
	"fmt"
)

type Logger struct{}

func New() (*Logger, error) {
	return &Logger{}, nil
}

func (l *Logger) Info(text string) {
	fmt.Print(text)
}

func (l *Logger) Error(error string) {
	fmt.Print(error)
}

func (l *Logger) Warn(text string) {
	fmt.Print(text)
}
