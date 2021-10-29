package myerr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// MyError ...
type MyError struct {
	Code     int
	Messages []string
	Stack    string
}

func (e *MyError) Error() string {
	l := len(e.Messages)
	return e.Messages[l-1]
}

func (e *MyError) Marshal() string {
	s, _ := json.Marshal(e)
	return string(s)
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)

		funcName := runtime.FuncForPC(pc).Name()
		arr := strings.Split(funcName, "/")
		name := ""
		if len(arr) > 0 {
			name = arr[len(arr)-1]
		} else {
			name = funcName
		}
		fmt.Fprintf(buf, "\t%s: %d\n", name, line)
	}
	return buf.Bytes()
}

func New(err error) *MyError {
	ne, ok := err.(*MyError)
	if ok {
		return ne
	}
	e := &MyError{
		Code:     http.StatusInternalServerError,
		Messages: make([]string, 0),
		Stack:    string(stack(2)),
	}
	e.Messages = append(e.Messages, fmt.Sprintf("%s", err))
	return e

}

func Newf(format string, args ...interface{}) *MyError {
	e := &MyError{
		Code:     http.StatusInternalServerError,
		Messages: make([]string, 0),
		Stack:    string(stack(2)),
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
