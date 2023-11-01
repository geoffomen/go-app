// 用于声明本模块所需的接口, 具体实现将在运行时注入
package httpserverimp

type Options struct {
	Port   int
	Logger LoggerIface
}

type LoggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Loglevel = int

const (
	LogLevelError = iota
	LogLevelWarnning
	LogLevelInfo
	LogLevelDebug
)

type MyErrorIface interface {
	Error() string
	Marshal() string
	GetCode() int
	GetLogLevel() Loglevel
}
