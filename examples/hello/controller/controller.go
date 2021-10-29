package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/geoffomen/go-app/examples/hello"
	"github.com/geoffomen/go-app/examples/hello/impl"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/vo"
)

// Controller ..
func Controller() map[string]interface{} {
	srv := impl.New(
		database.GetClient(),
	)

	return map[string]interface{}{

		"GET /exam/hello": func(sessInfo vo.SessionInfo) (interface{}, error) {
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
			fmt.Printf("%v", r)
			return nil, nil
		},

		"/exam/multipart": func(r *multipart.Form) (interface{}, error) {
			fmt.Printf("%v", r)
			return nil, nil
		},
	}
}
