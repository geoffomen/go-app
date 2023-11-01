package httpserverimp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/internal/app/common/base/vo"
)

func responseHandler() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		err := ctx.Next()
		if err != nil {
			writeError(ctx.responseWriter, err)
			return err
		}
		responses := ctx.getHandlerReflectValue().Call(ctx.getArgs())
		ctx.setResponse(responses)
		e := responses[1].Interface()
		if e != nil {
			writeError(ctx.responseWriter, e)
			return e.(error)
		} else {
			switch rt := responses[0].Interface().(type) {
			case http.ResponseWriter:
				// 在应用内已处理，这里不需重复处理
			default:
				writeSuccess(ctx.responseWriter, rt)
			}
		}
		return nil
	}
}

func writeSuccess(w http.ResponseWriter, obj any) {
	rt := vo.BaseRspDto{Code: 0, Msg: "ok", Data: obj}
	b, err := json.Marshal(rt)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", b)
	}
}

func writeError(w http.ResponseWriter, e interface{}) {
	if ne, ok := e.(MyErrorIface); ok {
		rt := vo.BaseRspDto{Code: ne.GetCode(), Msg: ne.Error()}
		b, _ := json.Marshal(rt)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(ne.GetCode())
		fmt.Fprintf(w, "%s", b)
	} else {
		var msg string
		if ne, ok := e.(error); ok {
			msg = ne.Error()
		} else {
			msg = "服务异常"
		}
		rt := vo.BaseRspDto{Code: http.StatusInternalServerError, Msg: msg}
		b, _ := json.Marshal(rt)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(rt.Code)
		fmt.Fprintf(w, "%s", b)
	}
}
