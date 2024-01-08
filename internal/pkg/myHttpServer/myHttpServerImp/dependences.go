// 用于声明本模块所需的接口, 具体实现将在运行时注入
package myHttpServerImp

import (
	"net/http"

	"ibingli.com/internal/pkg/myHttpServer"
)

type Options struct {
	Port        int
	Logger      loggerIface
	AuthHandler authIface
}

type loggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type authIface interface {
	Auth(req *http.Request) (*myHttpServer.SessionInfo, bool)
}

type Loglevel = int

const (
	LogLevelError = iota
	LogLevelWarnning
	LogLevelInfo
	LogLevelDebug
)

type myErrorIface interface {
	Error() string
	Marshal() string
	GetCode() int
	GetLogLevel() Loglevel
}
