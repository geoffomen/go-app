package httpserverimp

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var hdMap = make(map[string]http.HandlerFunc, 0)

type MyHttpServer struct {
	srv         *http.Server
	logger      LoggerIface
	middleWares []func(ctx *Ctx) error
}

func New(opts *Options) *MyHttpServer {
	srv := &MyHttpServer{
		srv: &http.Server{
			Addr: fmt.Sprintf(":%d", opts.Port),
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mp := fmt.Sprintf("%s%s", r.Method, r.URL.Path)
				if h, ok := hdMap[mp]; ok {
					h(w, r)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintln(w, "请求路径无效")
				}
			}),
		},
		logger: opts.Logger,
	}
	// 中间件的顺序很重要
	srv.addMiddleWare(logHandler())
	srv.addMiddleWare(responseHandler())
	srv.addMiddleWare(authHandler())
	srv.addMiddleWare(bindArgs())
	return srv
}

func (s *MyHttpServer) Listen() error {
	return s.srv.ListenAndServe()
}

func (s *MyHttpServer) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}

func (s *MyHttpServer) AddRouter(hs map[string]interface{}) error {
	for methodAndPath, handler := range hs {
		if methodAndPath == "" || handler == nil {
			panic(fmt.Sprintf("failed to register handler，url: %s, handler: %v", methodAndPath, handler))
		}

		var method string
		var reqPath string
		mNp := strings.Split(methodAndPath, " ")
		switch len(mNp) {
		case 1:
			method = "GET"
			reqPath = strings.Trim(mNp[0], " \t")
		case 2:
			method = strings.ToUpper(strings.Trim(mNp[0], " \t"))
			reqPath = strings.Trim(mNp[1], " \t")
		default:
			return fmt.Errorf("failed to register handler, expected format: [method path]，actual format: %s", methodAndPath)
		}
		mp := fmt.Sprintf("%s%s", method, reqPath)
		if _, ok := hdMap[mp]; !ok {
			hdMap[mp] = s.wrapper(handler)
			s.logger.Infof("handler注册成功，%s %s", method, reqPath)
		} else {
			return fmt.Errorf("conflict: %s already registed", mp)
		}
	}
	return nil
}

func (s *MyHttpServer) addMiddleWare(mdw func(ctx *Ctx) error) error {
	s.middleWares = append(s.middleWares, mdw)
	return nil
}

func (s *MyHttpServer) wrapper(handler interface{}) http.HandlerFunc {
	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	if handlerType.Kind() != reflect.Func {
		panic("failed to register handler, handler not a function")
	}
	if handlerType.NumOut() != 2 {
		panic("failed to register handler, expect two params, no more or less.")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := newCtx(w, r)
		ctx.setLogger(s.logger)
		ctx.setMiddleware(s.middleWares)
		ctx.setHandlerReflectType(handlerType)
		ctx.setHandlerReflectValue(handlerValue)
		ctx.Next()
	}
}
