package myLog

type Iface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Loglevel = int

const (
	LogLevelError = iota
	LogLevelWarnning
	LogLevelInfo
	LogLevelDebug
)
