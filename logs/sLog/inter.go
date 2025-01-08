package sLog

type InterLogger interface {
	SetLevel(string)

	// template, ...args
	Fatalf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Warnf(string, ...interface{})
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
}

type InterStringLog interface {
	String() string
	Log()
}

type InterStringLogName interface {
	String(string) string
	Log(string)
}
