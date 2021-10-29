package ginimp

import (
	"fmt"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/geoffomen/go-app/pkg/vo"

	"github.com/gin-gonic/gin"
)

// Gin ...
type Gin struct {
	engine *gin.Engine
	conf   *Configuration
}

type GinLogger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Configuration struct {
	Profile string
	Port    string
}

var (
	pathToHandler = make(map[string]interface{})
	pathToMethod  = make(map[string][]string)
	lgg           GinLogger
)

// NewGin ..
func New(conf Configuration, lg GinLogger) (*Gin, error) {
	if conf.Profile == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	lgg = lg
	// Creates a router without any middleware by default
	engine := gin.New()

	engine.Use(recovery())
	engine.Use(logger())
	engine.Use(responseHandler())
	engine.Use(httpMethodValidator())
	engine.Use(rateLimit())
	engine.Use(enableCors())
	engine.Use(authorization())
	engine.Use(bindQuery())
	engine.Use(bindJson())
	engine.Use(bindMutipartFormData())
	engine.Use(bindFromUrlencoded())

	return &Gin{
		engine: engine,
		conf:   &conf,
	}, nil

}

// RegisterHandler ...
func (s *Gin) RegisterHandler(m map[string]interface{}) {
	for methodAndPath, handler := range m {
		if methodAndPath == "" || handler == nil {
			panic(fmt.Sprintf("failed to register handler，url: %s, handler: %v", methodAndPath, handler))
		}

		var method string
		var reqPath string
		mNp := strings.Split(methodAndPath, " ")
		switch len(mNp) {
		case 1:
			reqPath = strings.Trim(mNp[0], " \t")
		case 2:
			method = strings.ToUpper(strings.Trim(mNp[0], " \t"))
			reqPath = mNp[1]
		default:
			panic(fmt.Sprintf("failed to register handler, expected format: [method path]，actual format: %s", methodAndPath))
		}

		reqPath = path.Join(reqPath)
		pathToHandler[method+reqPath] = handler
		if ms, ok := pathToMethod[reqPath]; ok {
			ms = append(ms, method)
			pathToMethod[reqPath] = ms
		} else {
			ms = make([]string, 0)
			ms = append(ms, method)
			pathToMethod[reqPath] = ms
		}

		switch method {
		case http.MethodGet:
			s.engine.GET(reqPath, wrapper(handler))
		case http.MethodHead:
			s.engine.HEAD(reqPath, wrapper(handler))
		case http.MethodPost:
			s.engine.POST(reqPath, wrapper(handler))
		case http.MethodPut:
			s.engine.PUT(reqPath, wrapper(handler))
		case http.MethodPatch:
			s.engine.PATCH(reqPath, wrapper(handler))
		case http.MethodDelete:
			s.engine.DELETE(reqPath, wrapper(handler))
		case http.MethodConnect:
			s.engine.Handle(method, reqPath, wrapper(handler))
		case http.MethodOptions:
			s.engine.OPTIONS(reqPath, wrapper(handler))
		case http.MethodTrace:
			s.engine.Handle(method, reqPath, wrapper(handler))
		case "":
			s.engine.GET(reqPath, wrapper(handler))
		default:
			panic(fmt.Sprintf("failed to register handler, not supported http method: %s", method))
		}
	}
}

// Start ..
func (s *Gin) Start() error {
	s.engine.Run(":" + s.conf.Port)
	return nil
}

func wrapper(handler interface{}) gin.HandlerFunc {
	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	if handlerType.Kind() != reflect.Func {
		panic("failed to register handler, handler not a function")
	}
	if handlerType.NumOut() != 2 {
		panic("failed to register handler, expect two params, no more or less.")
	}

	return func(c *gin.Context) {
		var args []reflect.Value
		arg, isExist := c.Get("args")
		if !isExist {
			switch numIn := handlerType.NumIn(); numIn {
			case 0:
				args = make([]reflect.Value, 0)
			case 1:
				fType := handlerType.In(0)
				if fType.String() == reflect.TypeOf(vo.SessionInfo{}).String() {
					sessionInfo, exist := c.Get("sessionInfo")
					if !exist {
						c.Error(fmt.Errorf("no session info"))
						c.Status(500)
						c.Abort()
						return
					}
					sessInfo, ok := sessionInfo.(vo.SessionInfo)
					if !ok {
						c.Error(fmt.Errorf("no session info"))
						c.Status(500)
						c.Abort()
						return
					}
					args = make([]reflect.Value, 1)
					args[0] = reflect.ValueOf(sessInfo)
				} else {
					c.Error(fmt.Errorf("require %d args, but no args in context", numIn))
					c.Status(500)
					c.Abort()
					return
				}
			default:
				c.Error(fmt.Errorf("require %d args, but no args in context", numIn))
				c.Status(500)
				c.Abort()
				return
			}
		} else {
			var ok bool
			args, ok = arg.([]reflect.Value)
			if !ok {
				c.Error(fmt.Errorf("args type error"))
				c.Status(500)
				c.Abort()
				return
			}
		}
		rsp := handlerValue.Call(args)
		c.Set("responses", rsp)
	}
}
