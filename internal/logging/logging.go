package logging

import (
	"log"
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type defaultLogger struct {
	debug *log.Logger
	info  *log.Logger
	error *log.Logger
}

// ...logger implementation
