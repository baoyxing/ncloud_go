package core

import "fmt"

const (
	LogFatal = iota
	LogError
	LogWarn
	LogInfo
)

type Logger interface {
	Log(level int, message ...interface{})
}

type DefaultLogger struct {
	Level int
}

func NewDefaultLogger(level int) *DefaultLogger {
	return &DefaultLogger{level}
}

func (logger DefaultLogger) Log(level int, message ...interface{}) {
	if level <= logger.Level {
		fmt.Println(message...)
	}
}

type DummyLogger struct {
}

func NewDummyLogger() *DummyLogger {
	return &DummyLogger{}
}

func (logger DummyLogger) Log(level int, message ...interface{}) {
}
