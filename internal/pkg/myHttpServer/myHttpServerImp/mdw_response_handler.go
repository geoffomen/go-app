package myHttpServerImp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"ibingli.com/internal/pkg/myHttpServer"
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
			case *myHttpServer.FileContentRspDto:
				return serveFileContent(ctx.responseWriter, ctx.request, rt.FileAbsPath)
			case myHttpServer.FileContentRspDto:
				return serveFileContent(ctx.responseWriter, ctx.request, rt.FileAbsPath)
			case *myHttpServer.AttachmentContentRspDto:
				return serveAttachmentContent(ctx.responseWriter, ctx.request, rt)
			case myHttpServer.AttachmentContentRspDto:
				return serveAttachmentContent(ctx.responseWriter, ctx.request, &rt)
			case *myHttpServer.InlineContentRspDto:
				return serveInlineContent(ctx.responseWriter, ctx.request, rt)
			case myHttpServer.InlineContentRspDto:
				return serveInlineContent(ctx.responseWriter, ctx.request, &rt)

			default:
				writeSuccess(ctx.responseWriter, rt)
			}
		}
		return nil
	}
}

func writeSuccess(w http.ResponseWriter, obj any) {
	rt := myHttpServer.BaseRspDto{Code: 0, Msg: "ok", Data: obj}
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
	if ne, ok := e.(myErrorIface); ok {
		rt := myHttpServer.BaseRspDto{Code: ne.GetCode(), Msg: ne.Error()}
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
		rt := myHttpServer.BaseRspDto{Code: http.StatusInternalServerError, Msg: msg}
		b, _ := json.Marshal(rt)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(rt.Code)
		fmt.Fprintf(w, "%s", b)
	}
}

func serveFileContent(w http.ResponseWriter, req *http.Request, absPath string) error {
	f, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer f.Close()
	// fi, err := os.Stat(absPath)
	// if err != nil {
	// 	return err
	// }
	// mt := fi.ModTime()
	mt := time.Now()
	http.ServeContent(w, req, f.Name(), mt, f)
	return nil
}

func serveAttachmentContent(w http.ResponseWriter, req *http.Request, data *myHttpServer.AttachmentContentRspDto) error {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", data.Name))
	http.ServeContent(w, req, data.Name, time.Now(), bytes.NewReader(data.Content))
	return nil
}

func serveInlineContent(w http.ResponseWriter, req *http.Request, data *myHttpServer.InlineContentRspDto) error {
	http.ServeContent(w, req, "", time.Now(), bytes.NewReader(data.Content))
	return nil
}
