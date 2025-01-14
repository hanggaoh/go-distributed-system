package util

import (
	"log"
	"os"
)

// Logger is a simple logger that writes to stdout.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new Logger.
func NewLogger(prefix string) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}
