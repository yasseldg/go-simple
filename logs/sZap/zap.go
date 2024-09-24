package sZap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	level zap.AtomicLevel
	*zap.SugaredLogger
}

func New(timeFormat, level string) (Logger, func() error) {

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

	l.SugaredLogger = zap.New(core).Sugar()

	l.SetLevel(level)

	return l, l.SugaredLogger.Sync
}

func (l Logger) SetLevel(level string) {
	err := l.level.UnmarshalText([]byte(level))
	if err != nil {
		// define default level as debug level
		l.level.SetLevel(zapcore.DebugLevel)
	}

	l.SugaredLogger.Infof("Logger -- Zap Sugar -- set level at ( %s )", l.level.String())
}
