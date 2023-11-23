package httpserverimp

import (
	"net/http"
	"reflect"

	"example.com/internal/app/common/base/vo"
)

// Ctx 一次http请求中的上下文信息
type Ctx struct {
	logger                 LoggerIface
	request                *http.Request
	responseWriter         http.ResponseWriter
	middleWareFuns         []func(ctx *Ctx) error
	currentMiddleWareIndex int
	args                   []reflect.Value
	response               []reflect.Value
	sessionInfo            *vo.SessionInfo
	handlerReflectType     reflect.Type
	handlerReflectValue    reflect.Value
}

func newCtx(w http.ResponseWriter, req *http.Request) *Ctx {
	return &Ctx{
		request:                req,
		responseWriter:         w,
		currentMiddleWareIndex: -1,
	}
}

func (c *Ctx) Next() error {
	c.currentMiddleWareIndex++
	var err error = nil
	if c.currentMiddleWareIndex < len(c.middleWareFuns) {
		err = c.middleWareFuns[c.currentMiddleWareIndex](c)
	}
	return err
}

func (c *Ctx) setMiddleware(mdws []func(ctx *Ctx) error) {
	c.middleWareFuns = mdws
}

func (c *Ctx) setLogger(loger LoggerIface) {
	c.logger = loger
}

func (c *Ctx) setHandlerReflectType(hft reflect.Type) {
	c.handlerReflectType = hft
}

func (c *Ctx) getHandlerReflectType() reflect.Type {
	return c.handlerReflectType
}

func (c *Ctx) setHandlerReflectValue(hfv reflect.Value) {
	c.handlerReflectValue = hfv
}

func (c *Ctx) getHandlerReflectValue() reflect.Value {
	return c.handlerReflectValue
}

func (c *Ctx) setArgs(args []reflect.Value) {
	c.args = args
}

func (c *Ctx) getArgs() []reflect.Value {
	return c.args
}

func (c *Ctx) setSessionInfo(si *vo.SessionInfo) {
	c.sessionInfo = si
}

func (c *Ctx) getSessionInfo() *vo.SessionInfo {
	return c.sessionInfo
}

func (c *Ctx) setResponse(response []reflect.Value) {
	c.response = response
}

func (c *Ctx) getResponnse() []reflect.Value {
	return c.response
}