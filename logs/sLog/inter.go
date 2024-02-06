package sLog

import sZap "github.com/yasseldg/go-simple/logs/sZap"

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

var logger Logger

func SetLogger(l Logger) {
	logger = l
}

func SetByName(name Name, level Level, timeformat string) func() error {
	switch name {
	default:
		zap, sync := sZap.New(timeformat, string(level))
		logger = zap
		return sync
	}
}

func Fatal(template string, args ...interface{}) {
	logger.Fatal(template, args...)
}

func Error(template string, args ...interface{}) {
	logger.Error(template, args...)
}

func Panic(template string, args ...interface{}) {
	logger.Panic(template, args...)
}

func Warn(template string, args ...interface{}) {
	logger.Warn(template, args...)
}

func Info(template string, args ...interface{}) {
	logger.Info(template, args...)
}

func Debug(template string, args ...interface{}) {
	logger.Debug(template, args...)
}
