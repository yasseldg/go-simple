package sZap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	level zap.AtomicLevel
	*zap.SugaredLogger
	encoderConfig zapcore.EncoderConfig
}

func New(timeFormat, level string) (*Logger, func() error) {

	l := &Logger{
		level: zap.NewAtomicLevel(),
	}

	l.encoderConfig = zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout(timeFormat),
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(l.encoderConfig), os.Stdout, l.level)
	l.SugaredLogger = zap.New(core).Sugar()

	l.SetLevel(level)

	return l, l.SugaredLogger.Sync
}

func (l *Logger) SetLevel(level string) {
	err := l.level.UnmarshalText([]byte(level))
	if err != nil {
		// define default level as debug level
		l.level.SetLevel(zapcore.DebugLevel)
	}

	l.SugaredLogger.Infof("Logger -- Zap Sugar -- set level at ( %s )", l.level.String())
}

func (l *Logger) SetWriteSyncer(callback func(string)) {
	if callback != nil {
		ws := &writeSyncer{callback: callback}
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(l.encoderConfig),
			zapcore.AddSync(ws),
			l.level,
		)
		l.SugaredLogger = zap.New(core).Sugar()
	}
}

// WriteSyncer is a zapcore.WriteSyncer implementation that calls a callback function with the log message.
// It implements the zapcore.WriteSyncer interface, allowing it to be used as a destination for zap logs.
// The callback function is called with the log message as a string.
// This is useful for logging to a custom output stream or for sending logs to a remote service.
// The WriteSyncer can be used with zapcore.NewCore to create a new zapcore.Core that writes logs to the callback function.
// The WriteSyncer can be used with zapcore.NewCore to create a new zapcore.Core that writes logs to the callback function.

type writeSyncer struct {
	callback func(string)
}

func (ws *writeSyncer) Write(p []byte) (n int, err error) {
	ws.callback(string(p))
	return len(p), nil
}

func (ws *writeSyncer) Sync() error {
	return nil
}
