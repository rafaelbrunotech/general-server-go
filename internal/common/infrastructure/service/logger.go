package service

import (
	"fmt"
)

type Logger struct{}

func NewLogger() (*Logger, error) {
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
