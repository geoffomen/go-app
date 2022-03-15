package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/webfw"

	"github.com/geoffomen/go-app/examples/hello"
	"github.com/geoffomen/go-app/examples/hello/impl"
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

		"/exam/error": func() (interface{}, error) {
			return srv.Error()
		},

		"/exam/req-rsp": func(req http.Request, rsp http.ResponseWriter) (interface{}, error) {
			fmt.Printf("%v %v", req, rsp)
			return nil, nil
		},

		"/exam/ioreader": func(r io.ReadCloser) (interface{}, error) {
			defer r.Close()
			b, err := io.ReadAll(r)
			fmt.Printf("%s", b)
			return nil, err
		},

		"/exam/multipart": func(r *multipart.Form) (interface{}, error) {
			fmt.Printf("%v", r)
			return nil, nil
		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
