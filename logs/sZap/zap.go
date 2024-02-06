package sZap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const EncodeTimeFormat = "2006.01.02 15:04:05"

type Logger struct {
	level  zap.AtomicLevel
	logger *zap.SugaredLogger
}

func New(timeFormat, level string) (Logger, func() error) {
	if timeFormat == "" {
		timeFormat = EncodeTimeFormat
	}

	l := Logger{
		level: zap.NewAtomicLevel(),
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout(timeFormat),
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, l.level)
	l.logger = zap.New(core).Sugar()

	l.SetLevel(level)

	return l, l.logger.Sync
}

func (l Logger) SetLevel(level string) {
	err := l.level.UnmarshalText([]byte(level))
	if err != nil {
		// define default level as debug level
		l.level.SetLevel(zapcore.DebugLevel)
	}

	l.logger.Infof("Logger -- Zap Sugar -- set level at ( %s )", l.level.String())
}

func (l Logger) Fatal(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l Logger) Error(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l Logger) Panic(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l Logger) Warn(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l Logger) Info(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l Logger) Debug(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}
