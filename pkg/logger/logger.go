package logger

import "fmt"

// Интерфейс логгера
type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

// Простая реализация логгера
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

func (l *SimpleLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
