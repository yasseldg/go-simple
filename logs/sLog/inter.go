package sLog

import (
	"sync"

	sZap "github.com/yasseldg/go-simple/logs/sZap"
)

type Logger interface {
	SetLevel(string)

	Fatal(template string, args ...interface{})
	Error(template string, args ...interface{})
	Panic(template string, args ...interface{})
	Warn(template string, args ...interface{})
	Info(template string, args ...interface{})
	Debug(template string, args ...interface{})
}

type Name string
type Level string

const (
	Zap = Name("zap")

	LevelFatal = Level("fatal")
	LevelError = Level("error")
	LevelPanic = Level("panic")
	LevelWarn  = Level("warn")
	LevelInfo  = Level("info")
	LevelDebug = Level("debug")
)

var (
	_onceCandle sync.Once
	_logger     Logger
)

func SetLogger(l Logger) {
	_onceCandle.Do(func() {
		_logger = l
	})
}

func SetByName(name Name, level Level, timeformat string) func() error {
	switch name {
	default:
		zap, clean := sZap.New(timeformat, string(level))
		SetLogger(zap)
		return clean
	}
}

func Fatal(template string, args ...interface{}) {
	_logger.Fatal(template, args...)
}

func Error(template string, args ...interface{}) {
	_logger.Error(template, args...)
}

func Panic(template string, args ...interface{}) {
	_logger.Panic(template, args...)
}

func Warn(template string, args ...interface{}) {
	_logger.Warn(template, args...)
}

func Info(template string, args ...interface{}) {
	_logger.Info(template, args...)
}

func Debug(template string, args ...interface{}) {
	_logger.Debug(template, args...)
}
