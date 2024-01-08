package myErrImp

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
)

type Loglevel = int

const (
	LogLevelError = iota
	LogLevelWarnning
	LogLevelInfo
	LogLevelDebug
)

// MyError ...
type MyError struct {
	Code     int      // 错误码
	Messages []string // 额外的附加信息
	Stack    string   // 当前错误链首个错误发生时的调用栈信息
	LogLevel Loglevel // 当前错误链应当使用的日志等级（如果记录日志的话），默认为LogLevelError
}

func (e *MyError) Error() string {
	return e.getLastMessage()
}

func (e *MyError) Marshal() string {
	return fmt.Sprintf("\ncode: %d. \nMessages: %s. \nStack: \n%s", e.Code, strings.Join(e.Messages, "; "), e.Stack)
}

func New(err error, additionalMsg ...string) *MyError {
	ne, ok := err.(*MyError)
	if ok {
		ne.Messages = append(ne.Messages, additionalMsg...)
		return ne
	}
	e := &MyError{
		Code:     http.StatusInternalServerError,
		Messages: make([]string, 0),
		Stack:    string(debug.Stack()),
	}
	e.Messages = append(e.Messages, err.Error())
	e.Messages = append(e.Messages, additionalMsg...)
	return e

}

func Newf(format string, args ...interface{}) *MyError {
	e := &MyError{
		Code:     http.StatusInternalServerError,
		Messages: make([]string, 0),
		Stack:    string(debug.Stack()),
	}
	e.Messages = append(e.Messages, fmt.Sprintf(format, args...))
	return e
}

func NewfWithCode(code int, format string, args ...interface{}) *MyError {
	e := Newf(format, args...)
	e.Code = code
	return e

}

func (e *MyError) AddMsgf(format string, args ...interface{}) *MyError {
	e.Messages = append(e.Messages, fmt.Sprintf(format, args...))
	return e
}

func (e *MyError) SetCode(code int) *MyError {
	e.Code = code
	return e
}

func (e *MyError) GetCode() int {
	return e.Code
}

func (e *MyError) SetLogLevel(logLevel int) *MyError {
	e.LogLevel = logLevel
	return e
}

func (e *MyError) GetLogLevel() int {
	return e.LogLevel
}

func (e *MyError) getLastMessage() string {
	l := len(e.Messages)
	if l > 0 {
		return e.Messages[l-1]
	}
	return ""
}
