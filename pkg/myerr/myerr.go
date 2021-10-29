package myerr


import (
	"fmt"
	"net/http"
	"runtime/debug"
)

const (
	LogLevelDebug int = iota
	LogLevelWarnning
	LogLevelInfo
)

// MyError ...
type MyError struct {
	Code     int
	Messages []string
	Stack    string
	LogLevel int // 决定如何记录该错误，默认为LogLevelDebug，表示应该记录在错误日志，其它等级则不建议记录在错误日志。需要显式指定
}

func (e *MyError) Error() string {
	l := len(e.Messages)
	return e.Messages[l-1]
}

func (e *MyError) Marshal() string {
	return fmt.Sprintf("code: %d, Messages: %s, Stack: %s", e.Code, e.Messages, e.Stack)
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
	e.Messages = append(e.Messages, fmt.Sprintf("%s", err))
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

func (e *MyError) SetCode(code int) *MyError {
	e.Code = code
	return e
}

func (e *MyError) AddMsgf(format string, args ...interface{}) *MyError {
	e.Messages = append(e.Messages, fmt.Sprintf(format, args...))
	return e
}

func (e *MyError) SetLogLevel(logLevel int) *MyError {
	e.LogLevel = logLevel
	return e
}

func (e *MyError) GetLastMessage() string {
	l := len(e.Messages)
	if l > 0 {
		return e.Messages[l-1]
	}
	return ""
}
