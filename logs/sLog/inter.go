package sLog

import (
	"strings"
	"sync"

	sZap "github.com/yasseldg/go-simple/logs/sZap"
)

type Logger interface {
	SetLevel(string)

	Fatalf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Debugf(template string, args ...interface{})
}

type Name string
type Level string

const (
	EncodeTimeFormat = "2006.01.02 15:04:05"

	Zap = Name("zap")

	LevelFatal = Level("fatal")
	LevelError = Level("error")
	LevelPanic = Level("panic")
	LevelWarn  = Level("warn")
	LevelInfo  = Level("info")
	LevelDebug = Level("debug")
)

var (
	_once   sync.Once
	_logger Logger

	_indentation int
)

func SetLogger(l Logger) {
	_once.Do(func() {
		_logger = l
	})
}

func SetByName(name Name, level Level, timeformat string) func() error {

	if len(timeformat) == 0 {
		timeformat = EncodeTimeFormat
	}

	_indentation = len(timeformat) + 2

	switch name {
	default:
		_indentation += 10

		zap, clean := sZap.New(timeformat, string(level))
		SetLogger(zap)
		return clean
	}
}

func SetLevel(level Level) {
	_logger.SetLevel(string(level))
}

func Fatal(template string, args ...interface{}) {
	_logger.Fatalf(template, args...)
}

func Error(template string, args ...interface{}) {
	_logger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	_logger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	_logger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	_logger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	_logger.Debugf(template, args...)
}

// Lines get a message with a new line for each %l in the template
// use 1 blank space after %l for alignment, or more than 1 for indentation.
func Lines(template string) string {
	if strings.Contains(template, "%l") {
		// Obtenemos el nivel de log para calcular la cantidad de espacios para alinear.
		wrappedMessage := strings.Replace(template, "%l", "\n"+strings.Repeat(" ", _indentation), -1)
		return wrappedMessage
	}
	return template
}
