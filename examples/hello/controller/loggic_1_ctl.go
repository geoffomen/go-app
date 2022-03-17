package controller

import (
	"fmt"
	"net/http"

	"github.com/storm-5/go-app/pkg/database"
	"github.com/storm-5/go-app/pkg/webfw"

	"github.com/storm-5/go-app/examples/hello"
	"github.com/storm-5/go-app/examples/hello/impl"
)

var (
	controllers map[string]interface{} = make(map[string]interface{})
	srv         *impl.Service
)

// Controller ..
func Controller() map[string]interface{} {
	srv = impl.New(
		database.GetClient(),
	)
	return controllers
}

func init() {
	m := map[string]interface{}{

		"GET /exam/hello": func(sessInfo webfw.SessionInfo) (interface{}, error) {
			return srv.SayHello(sessInfo)
		},

		"POST /exam/echo": func(args hello.EchoReqDto) (interface{}, error) {
			return srv.Echo(args)
		},

		"/exam/req-rsp": func(req http.Request, rsp http.ResponseWriter) (interface{}, error) {
			fmt.Printf("%v %v", req, rsp)
			return nil, nil
		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
