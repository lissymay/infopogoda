package logger

import "fmt"

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string, err error)
}

type SimpleLogger struct{}

func New() *SimpleLogger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func (l *SimpleLogger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}

func (l *SimpleLogger) Error(msg string, err error) {
	fmt.Printf("[ERROR] %s: %v\n", msg, err)
}
