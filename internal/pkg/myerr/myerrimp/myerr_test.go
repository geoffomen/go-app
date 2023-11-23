package myerrimp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMyerr(t *testing.T) {
	var n interface{} = nil
	v, ok := n.(MyError)
	if ok {
		fmt.Println(v)
	}
	err := Newf("error occur: %s", "reason")
	err = New(err).SetCode(http.StatusNotFound).SetLogLevel(LogLevelInfo).AddMsgf("请求资源不存在")
	fmt.Printf("%s", err.Error())
	fmt.Printf("%s", err.Marshal())
}
